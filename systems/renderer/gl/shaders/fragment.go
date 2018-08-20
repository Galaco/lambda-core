package shaders

var Fragment = `
    #version 410

	// Interpolated values from the vertex shaders
	in vec2 UV;

    //out vec4 frag_colour;
    out vec3 frag_colour;

	uniform sampler2D baseTexture;

    void main() {
		// Output color = color of the texture at the specified UV
//		if (textureSize(baseTexture, 0).x > 0) {
 //   		frag_colour = texture( baseTexture, UV ).rgb;
//	    } else {
//        	frag_colour = vec3(1, 1, 1);
//		}
frag_colour = texture( baseTexture, UV ).rgb;
    }
` + "\x00"