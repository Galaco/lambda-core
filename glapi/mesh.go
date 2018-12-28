package glapi

import (
	opengl "github.com/go-gl/gl/v4.1-core/gl"
)

const Line = uint32(opengl.LINE)
const Triangles = uint32(opengl.TRIANGLES)

var vertexDrawMode = Line

func SetLineWidth(width float32) {
	opengl.LineWidth(width)
}

func SetVertexDrawMode(drawMode uint32) {
	vertexDrawMode = drawMode
}


func DrawArray(offset int, length int) {
	opengl.DrawArrays(vertexDrawMode, int32(offset), int32(length))
}