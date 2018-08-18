package shaders

var Fragment = `
    #version 410

	// Interpolated values from the vertex shaders
	in vec2 UV;

    out vec4 frag_colour;
    //out vec3 frag_colour;

	uniform sampler2D baseTexture;

    void main() {
        frag_colour = vec4(1, 1, 1, 0.025);

		// Output color = color of the texture at the specified UV
    	//frag_colour = texture( baseTexture, UV ).rgb;
    }
` + "\x00"