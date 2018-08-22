# -Go-me-Engine
A random game engine in go


# What can this do?
Right now, this project can load any standard v20 BSP for any Source Engine game, and display brush faces textured
with vtf low-resolution data, obtained from that games multi-part vpk.

See `main.go` for now to specify the .bsp and .vpks containing textures.


# What will this do?
Not sure yet. For now its just a bt of fun. Next goals are to to be able to load in .mdl meshes, and high-res vtf data,
as well as displacements. After that, most likely pakfile data extraction & skybox, and plenty of optimisation.