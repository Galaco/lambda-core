# Go Source Engine Demo
A sort-of game engine in go. The engine really just consists of a set of simple constructs for managing a game loop, event queue and in-memory game
resources. Beyond that, it's really up to a game to create any systems and components, and whatever else. Engine code is
located in `engine/`.

This project contains a layer built on top of these constructs to be able to work with data types from Valve's Source
Engine. Game specific components with their associated structs, and systems live in `components/` and `systems/` 
respectively. `valve/` contains loading and parsing code for Sourcing engine data, including some library wrappers.


# What can this do?
Right now, this project can load any standard (i.e. no game specific lump modification) v20 BSP for any Source Engine 
game, although it is tested against Counterstrike: Source official and community maps It can:
* Display all bsp faces
* Load high-resolution materials (both .vmt and .vtf parsing, but only baseTexture is used) from both game VPK and 
target .bsp pakfile
* Parse and load the entdata lump (for now draws a small primitive at their origin)
* Builds a complete bsp tree, and visibility data. Is currently a little buggy, but will update visible cluster faces
based on current camera position

See `main.go` for now to specify the .bsp and .vpks containing textures.


# What will this do?
Not sure yet. This is just for fun, I'll have to see just how complete I want this to be.

### Minimal run:
```
package main

import "github.com/galaco/go-me-engine/engine"

func main() {
	Application := engine.NewEngine()

	Application.Initialise()

	Application.Run()
}
```