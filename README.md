# Gource Engine
Gource is a game engine written in golang designed that loads Valve's Source Engine projects. Put simply, pointing this projects configuration at
a Source Engine game installation directory will allow for loading that targets .bsp maps and contents.


## Current state
You can build this right now, and, assuming you set the configuration to point to an existing Source game installation (this is tested against CS:S and CS:GO):
* Loads game data files from projects gameinfo.txt
* Load BSP map
* Load high-resolution texture data for bsp faces, including pakfile entries
* Full visibility data support

##### Counterstrike: Source de_dust2.bsp
![de_dust2](https://raw.githubusercontent.com/galaco/Gource-Engine/master/Documents/de_dust2.jpg)


## What will this do?
The end goal is to be able to point this application at a source engine game, with a bsp as the target, and be able to
load and play that map.


## Getting started
There is a small amount of configuration required to get this project running, beyond `dep ensure`.
* For best results, you need a source engine game installed already.
* Copy `config.example.json` to `config.json`, and update the `gameDirectory` property to point to whatever game installation
you are targeting (e.g. HL2 would be `<steam_dir>/steamapps/common/hl2`).

## Contributing
There is loads to do! Right now there are a few core issues that need fixing, and loads of fundamental features to add. Here
are just a few!
* StudioModel library needs finishing before props can be properly added
* No Physics
* A vulkan renderer would be a huge step forward, particularly this early on. Abstracting a mesh away from ogl would also help
* Displacement support incomplete - generation is buggy, and visibility checks cull displacements always
* Additional game support/testing in BSP library
