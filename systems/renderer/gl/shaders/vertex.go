package shaders

var Vertex = `
    #version 410

	uniform mat4 projection;
	uniform mat4 view;
	uniform mat4 model;

	layout(location = 1) in vec2 vertexUV;

	// Output data ; will be interpolated for each fragment.
	out vec2 UV;

    in vec3 vp;

    void main() {
        gl_Position = projection * view * model * vec4(vp, 1.0);

    	// UV of the vertex. No special space for this one.
    	//UV = vertexUV;
    }
` + "\x00"