package glapi

import (
	opengl "github.com/go-gl/gl/v4.1-core/gl"
)

const Front = opengl.FRONT
const Back = opengl.BACK
const FrontAndBack = opengl.FRONT_AND_BACK
const DepthTestLEqual = opengl.LEQUAL
const WindingClockwise = opengl.CW
const WindingCounterClockwise = opengl.CCW

const MaskColourBufferBit = opengl.COLOR_BUFFER_BIT
const MaskDepthBufferBit = opengl.DEPTH_BUFFER_BIT

func ClearColour(red float32, green float32, blue float32, alpha float32) {
	opengl.ClearColor(red, green, blue, alpha)
}

func Clear(bufferBits ...uint32) {
	mask := uint32(0)
	for _,buf := range bufferBits {
		mask = mask | buf
	}
	opengl.Clear(mask)
}

func EnableBlend() {
	opengl.Enable(opengl.BLEND)
	opengl.BlendFunc(opengl.SRC_ALPHA, opengl.ONE_MINUS_SRC_ALPHA)
}

func DisableBlend(){
	opengl.Disable(opengl.BLEND)
}

func EnableDepthTest() {
	opengl.Enable(opengl.DEPTH_TEST)
	opengl.DepthFunc(opengl.LEQUAL)
}

func DisableDepthTest() {
	opengl.Disable(opengl.DEPTH_TEST)
}

func EnableCullFace(cullSide uint32, winding uint32) {
	opengl.Enable(opengl.CULL_FACE)
	opengl.CullFace(cullSide)
	opengl.FrontFace(winding)
}