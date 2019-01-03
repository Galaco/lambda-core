package vk

import (
	"github.com/galaco/Gource-Engine/client/scene/world"
	"github.com/galaco/Gource-Engine/core/entity"
	"github.com/galaco/Gource-Engine/core/logger"
	"github.com/galaco/Gource-Engine/core/model"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/vulkan-go/glfw/v3.3/glfw"
	"github.com/vulkan-go/vulkan"
)

type Renderer struct {

}

func (manager *Renderer) Initialize() {
	vulkan.SetGetInstanceProcAddr(glfw.GetVulkanGetInstanceProcAddress())
	err := vulkan.Init()
	if err != nil {
		logger.Fatal(err)
	}
}

func (manager *Renderer) StartFrame(*entity.Camera) {

}

func (manager *Renderer) LoadShaders() {

}

func (manager *Renderer) DrawBsp(*world.World) {

}

func (manager *Renderer) DrawSkybox(*world.Sky) {

}

func (manager *Renderer) DrawModel(*model.Model, mgl32.Mat4) {

}

func (manager *Renderer) DrawSkyMaterial(model *model.Model) {

}

func (manager *Renderer) SetWireframeMode(bool) {

}

func (manager *Renderer) EndFrame() {

}

func (manager *Renderer) Unregister() {

}
