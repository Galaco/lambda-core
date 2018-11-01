package shaders

var Vertex = `
    #version 410

	uniform mat4 projection;
	uniform mat4 view;
	uniform mat4 model;

    layout(location = 0) in vec3 vertexPosition;
	layout(location = 1) in vec2 vertexUV;
	layout(location = 2) in vec2 vertexNormal;

	// Output data ; will be interpolated for each fragment.
	out vec2 UV;

    void main() {
        gl_Position = projection * view * model * vec4(vertexPosition, 1.0);

    	// UV of the vertex. No special space for this one.
    	UV = vertexUV;
    }
` + "\x00"
