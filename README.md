[![GoDoc](https://godoc.org/github.com/galaco/lambda-core?status.svg)](https://godoc.org/github.com/galaco/lambda-core)
[![Go report card](https://goreportcard.com/badge/github.com/galaco/lambda-core)](https://goreportcard.com/badge/github.com/galaco/lambda-core)
[![GolangCI](https://golangci.com/badges/github.com/galaco/lambda-core.svg)](https://golangci.com)
[![Build Status](https://travis-ci.com/Galaco/lambda-core.svg?branch=master)](https://travis-ci.com/Galaco/lambda-core)
[![codecov](https://codecov.io/gh/Galaco/lambda-core/branch/master/graph/badge.svg)](https://codecov.io/gh/Galaco/lambda-core)
[![CircleCI](https://circleci.com/gh/Galaco/lambda-core.svg?style=svg)](https://circleci.com/gh/Galaco/lambda-core)

# Lambda Core
Lambda Core provides a semi-comprehensive set of tools to build practically any Source Engine tool from. Any module can be used 
in isolation, but its recommended to utilise at least the FileSystem and ResourceManager modules if any loader is used.

##### See [https://github.com/galaco/Lambda-Client](https://github.com/galaco/Lambda-Client) for a working BSP renderer built on top of this library.

### Current features
* GameInfo.txt parser for existing games
* Full filesystem loader and searcher for all GameInfo defined paths, including pakfile and vpks
* Vmt and Vtf parsing
* Basic .mdl parsing (useable, incomplete)
* Full bsp loading utilities

## Contributing
There is loads to do! Right now there are a few core issues that need fixing, and loads of fundamental features to add. Here
are just a few!
* StudioModel library needs finishing before props can be properly added. There are some issues around multiple stripgroups per mesh, multiple
materials per prop, mdl data not fully loaded, and likely more
* Implement physics (probably bullet physics? Accurate VPhysics is probably not worthwhile, but needs investigation)
* Displacement support incomplete - generation is buggy, and visibility checks cull displacements always (visible when outside of world only)
* Additional game support/testing in BSP library



#### Some documentation

##### Entity
Provides an interface, and generic implementation of a game entity, as well as
3d transform struct and camera implementation.

##### Event
Event provides a very simple emitter/subscriber manager to allow engine event 
processing. It can be used for handling internal engine events, or for game logic,
although that isn't recommended.

##### Filesystem
Source Engine is a little annoying in that there are potentially unlimited possible
locations that engine resources can be located. Filesystem provides a way to register 
and organise any potential resource path or filesystem, while preserving filesystem type
search priority.

##### Loader
Loader generates data structures for Source formats from file streams. Materials, textures,
meshes etc.

##### Logger
Logger is a simple module to abstract out different print() priorities, and (eventually) locations 
than just stdout.

##### Material 
Material provides a more GPU friendly vmt material implementation

##### Mesh
Provides a set of common data formats for vertex data in the bsp or compiled
.mdl props.

##### Model
Models are combinations of simple data structures that represent a single higher 
level visual object. e.g. []Mesh is a studio model, []Face is bsp data etc.

##### Resource
Resource provides a management struct for tracking what game resources have been 
loaded from the filesystem. When a map is loaded, all found materials, textures, models
should be added to the ResourceManager, so they can be loaded only once, and cleaned up
correctly when no longer needed.

##### Scene
Provides a simple scene struct that contains bsp face and staticprop information

##### Texture
Provides a set of GPU friendly texture formats. For now OpenGL usage is enforced, but 
abstracting that out should be doable.