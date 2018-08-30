package gl

import (
	"fmt"
	"github.com/galaco/go-me-engine/systems/renderer/gl/shaders"
	opengl "github.com/go-gl/gl/v4.1-core/gl"
	"log"
	"strings"
)

type Context struct {
	context uint32
}

func (context *Context) AddShader(source string, shaderType uint32) {
	shader, err := context.compileShader(source, shaderType)
	if err != nil {
		panic(err)
	}

	opengl.AttachShader(context.context, shader)
}

func (context *Context) Finalize() {
	opengl.LinkProgram(context.context)
}

func (context *Context) UseProgram() {
	opengl.UseProgram(context.context)
}

func (context *Context) GetUniform(name string) int32 {
	return opengl.GetUniformLocation(context.context, opengl.Str(name+"\x00"))
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

func NewContext() Context {
	if err := opengl.Init(); err != nil {
		panic(err)
	}

	version := opengl.GoStr(opengl.GetString(opengl.VERSION))
	log.Println("OpenGL version", version)

	context := Context{
		context: opengl.CreateProgram(),
	}
	context.AddShader(shaders.Vertex, opengl.VERTEX_SHADER)
	context.AddShader(shaders.Fragment, opengl.FRAGMENT_SHADER)
	context.Finalize()

	return context
}
