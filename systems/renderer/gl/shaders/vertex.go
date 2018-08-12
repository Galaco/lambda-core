package shaders

var Vertex = `
    #version 410

	uniform mat4 projection;
	uniform mat4 view;
	uniform mat4 model;

    in vec3 vp;

    void main() {
        gl_Position = projection * view * model * vec4(vp, 1.0);
    }
` + "\x00"