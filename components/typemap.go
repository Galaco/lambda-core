package components

import "github.com/galaco/go-me-engine/engine/core"

// What is this? It provides a means to identify different components
// without using reflection, which although elegant is costly
// in performance

var T_CameraComponent = core.EType("CameraComponent")
var T_RenderableComponent = core.EType("RenderableComponent")
var T_BspComponent = core.EType("BspComponent")