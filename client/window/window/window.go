package window

import (
	"github.com/galaco/Gource-Engine/core/logger"
	"github.com/vulkan-go/glfw/v3.3/glfw"
	"runtime"
)

// Create constructs a new GLFW Window
func Create(width int, height int, name string) *glfw.Window {
	runtime.LockOSThread()
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, name, nil, nil)
	if err != nil {
		logger.Fatal(err)
	}

	window.MakeContextCurrent()

	return window
}
