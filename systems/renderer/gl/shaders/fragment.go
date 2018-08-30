package shaders

var Fragment = `
    #version 410

	in vec2 UV;

    out vec4 frag_colour;

	uniform sampler2D baseTexture;

    void main() {
		// Output color = color of the texture at the specified UV
		frag_colour = texture( baseTexture, UV ).rgba;
		//frag_colour = vec4(1, 1, 1, 0.05);
    }
` + "\x00"
