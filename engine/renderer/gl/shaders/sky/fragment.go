package sky

var Fragment = `
    #version 410

	in vec3 UV;

    out vec4 frag_colour;

	uniform samplerCube cubemapTexture;

    void main() {
		// Output color = color of the texture at the specified UV
		frag_colour = texture( cubemapTexture, UV );
    }
` + "\x00"
