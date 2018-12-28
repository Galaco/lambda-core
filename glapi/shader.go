package glapi

import (
	"fmt"
	opengl "github.com/go-gl/gl/v4.1-core/gl"
	"strings"
)

type ShaderType uint32

const VertexShader = ShaderType(opengl.VERTEX_SHADER)
const FragmentShader = ShaderType(opengl.FRAGMENT_SHADER)

type Context struct {
	context uint32
}

func (ctx *Context) Id() uint32 {
	return ctx.context
}

func (ctx *Context) AddShader(source string, shaderType ShaderType) error {
	shader, err := ctx.compileShader(source, shaderType)
	if err != nil {
		return err
	}

	opengl.AttachShader(ctx.context, shader)

	return nil
}

func (ctx *Context) Finalize() {
	opengl.LinkProgram(ctx.context)
}

func (ctx *Context) UseProgram() {
	opengl.UseProgram(ctx.context)
}

func (ctx *Context) GetUniform(name string) int32 {
	return opengl.GetUniformLocation(ctx.context, opengl.Str(name+"\x00"))
}

func (ctx *Context) compileShader(source string, shaderType ShaderType) (uint32, error) {
	shader := opengl.CreateShader(uint32(shaderType))

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

func (ctx *Context) Destroy() {
	opengl.DeleteShader(ctx.Id())
}

func NewShader() Context {
	if err := opengl.Init(); err != nil {
		panic(err)
	}

	context := Context{
		context: opengl.CreateProgram(),
	}

	return context
}

