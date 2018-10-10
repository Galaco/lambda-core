# Gource Engine
A sort-of game engine in go. The engine really just consists of a set of simple constructs for managing a game loop, event queue and in-memory game
resources. Beyond that, it's really up to a game to create any systems and components, and whatever else. Engine code is
located in `engine/`.

This project contains a layer built on top of these constructs to be able to work with data types from Valve's Source
Engine. Game specific components with their associated structs, and systems live in `components/` and `systems/` 
respectively. `valve/` contains loading and parsing code for Sourcing engine data, including some library wrappers.


## What can this do?
Right now, this project can load any standard (i.e. no game specific lump modification) v20 BSP for any Source Engine 
game, although it is tested against Counterstrike: Source official and community maps It can:
* Display all bsp faces
* Load high-resolution materials (both .vmt and .vtf parsing, but only baseTexture is used) from both game VPK and 
target .bsp pakfile
* Parse and load the entdata lump (for now draws a small primitive at their origin)
* Builds a complete bsp tree, and visibility data. Is currently a little buggy, but will update visible cluster faces
based on current camera position

See `main.go` for now to specify the .bsp and .vpks containing textures.


## What will this do?
The end goal is to be able to point this application at a source engine game, with a bsp as the target, and be able to
load and play that map.


## Getting started
There is a small amount of configuration required to get this project running, beyond `dep ensure`.
* Copy `config.example.json` to `config.json`, and update the `gameDirectory` property to point to whatever game installation
you are targeting (e.g. HL2 would be `<steamapps>/common/hl2`).

## Contributing
There is loads to do! Right now there are a few core issues that need fixing, and loads of fundamental features to add. Here
are just a few!
* StudioModel library needs finishing
* Visibility data sometimes culls wrong faces on certain clusters (de_dust2 exhibits this a lot)
* No VPhysics
* A vulkan renderer would be a huge step forward
* Displacement support
* Additional game support/testing in BSP library
* Multi-VPK support