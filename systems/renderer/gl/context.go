package gl

import (
	opengl "github.com/go-gl/gl/v2.1/gl"
	"strings"
	"fmt"
	"log"
	"github.com/galaco/go-me-engine/systems/renderer/gl/shaders"
)

type Context struct {
	context uint32
}


func (context *Context) AddShader(source string) {
	shader, err := context.compileShader(source, opengl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	opengl.AttachShader(context.context, shader)
}

func (context *Context) Finalize() {
	opengl.LinkProgram(context.context)
}

func (context *Context) compileShader(source string, shaderType uint32) (uint32, error) {
	shader := opengl.CreateShader(shaderType)

	csources, free := opengl.Strs(source)
	opengl.ShaderSource(shader, 1, csources, nil)
	free()
	opengl.CompileShader(shader)

	var status int32
	opengl.GetShaderiv(shader, opengl.COMPILE_STATUS, &status)
	if status == opengl.FALSE {
		var logLength int32
		opengl.GetShaderiv(shader, opengl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		opengl.GetShaderInfoLog(shader, logLength, nil, opengl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func NewContext() Context{
	if err := opengl.Init(); err != nil {
		panic(err)
	}

	version := opengl.GoStr(opengl.GetString(opengl.VERSION))
	log.Println("OpenGL version", version)

	context := Context{
		context: opengl.CreateProgram(),
	}
	context.AddShader(shaders.Vertex)
	context.AddShader(shaders.Fragment)
	context.Finalize()

	return context
}