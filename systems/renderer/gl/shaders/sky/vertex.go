package sky

var Vertex = `
    #version 410

	uniform mat4 projection;
	uniform mat4 view;
	uniform mat4 model;

    layout(location = 0) in vec3 vertexPosition;

	// Output data ; will be interpolated for each fragment.
	out vec3 UV;

    void main() {
		vec4 WVP_Pos = (projection * view * model) * vec4(vertexPosition, 1.0);
    	gl_Position = WVP_Pos.xyww;
    	UV = vertexPosition;
    }
` + "\x00"
