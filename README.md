[![GoDoc](https://godoc.org/github.com/Galaco/Lambda-Core?status.svg)](https://godoc.org/github.com/Galaco/Lambda-Core)
[![Go report card](https://goreportcard.com/badge/github.com/galaco/Lambda-Core)](https://goreportcard.com/badge/github.com/galaco/Lambda-Core)
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
