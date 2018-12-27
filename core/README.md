### Core

Core contained shared code for any implementing applications. It consists of a set of 
utilities for loading and managing a Source Engine games resources.

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