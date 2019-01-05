package gl

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

const frustumOutside = 0
const frustumIntersect = 1
const frustumInside = 2

type Frustum struct {
	cameraPosition mgl32.Vec3
	cameraReferentialX mgl32.Vec3
	cameraReferentialY mgl32.Vec3
	cameraReferentialZ mgl32.Vec3

	tang float32
	ratio float32
	nearD float32
	farD float32
	width float32
	height float32

	sphereFactorX float32
	sphereFactorY float32
}

func (frustum *Frustum) setCamInternals(angle float32, ratio float32, nearD float32, farD float32) {
	frustum.ratio = ratio
	frustum.nearD = nearD
	frustum.farD = farD

	frustum.tang = float32(math.Tan(mgl64.DegToRad(float64(angle) * 0.5)))
	frustum.height = nearD * frustum.tang
	frustum.width = frustum.height * ratio


	angle = mgl32.DegToRad(angle)
	// compute width and height of the near and far plane sections
	frustum.tang = float32(math.Tan(float64(angle)))
	frustum.sphereFactorY = float32(1.0/math.Cos(float64(angle)))

	// compute half of the the horizontal field of view and sphereFactorX
	anglex := math.Atan(float64(frustum.tang*ratio))
	frustum.sphereFactorX = float32(1.0/math.Cos(float64(anglex)))
}

func (frustum *Frustum) setCamDef(p mgl32.Vec3, l mgl32.Vec3, u mgl32.Vec3) {
	frustum.cameraPosition = p

	frustum.cameraReferentialZ = l.Sub(p).Normalize()

	frustum.cameraReferentialX = (mgl32.Vec3{
		frustum.cameraReferentialZ.X() * u.X(),
		frustum.cameraReferentialZ.Y() * u.Y(),
		frustum.cameraReferentialZ.Z() * u.Z(),
	}).Normalize()

	frustum.cameraReferentialY = (mgl32.Vec3{
		frustum.cameraReferentialX.X() * frustum.cameraReferentialZ.X(),
		frustum.cameraReferentialX.Y() * frustum.cameraReferentialZ.Z(),
		frustum.cameraReferentialX.Z() * frustum.cameraReferentialZ.Z(),
	}).Normalize()
}

func (frustum *Frustum) SphereInFrustum(p mgl32.Vec3, radius float32) int {
	var d float32
	var az,ax,ay float32
	var result = frustumInside
	v := p.Sub(frustum.cameraPosition)


	az = v.Dot(frustum.cameraReferentialZ.Mul(-1))
	if az > frustum.farD + radius || az < frustum.nearD-radius {
		return frustumOutside
	}

	if az > frustum.farD - radius || az < frustum.nearD+radius {
		result = frustumIntersect
	}

	ay = v.Dot(frustum.cameraReferentialY)
	d = frustum.sphereFactorY * radius
	az *= frustum.tang
	if (ay > az+d || ay < -az-d) {
		return frustumOutside
	}
	if (ay > az-d || ay < -az+d) {
		result = frustumIntersect
	}

	ax = v.Dot(frustum.cameraReferentialX)
	az *= frustum.ratio
	d = frustum.sphereFactorX * radius
	if ax > az+d || ax < -az-d {
		return frustumOutside
	}
	if ax > az-d || ax < -az+d {
		result = frustumIntersect
	}

	return result
}

func (frustum *Frustum) PointInFrustum(p mgl32.Vec3) int {
	return frustum.SphereInFrustum(p, 0.01)
}

func (frustum *Frustum) AABBInFrustum(mins mgl32.Vec3, maxs mgl32.Vec3) int {
	size := maxs.Sub(mins)
	radius := math.Abs(float64(size.Len())) / 2


	return frustum.SphereInFrustum(size.Mul(0.5), float32(radius))
}