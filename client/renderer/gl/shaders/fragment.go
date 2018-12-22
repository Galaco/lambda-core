package shaders

var Fragment = `
    #version 410

	uniform int useLightmap;

	uniform sampler2D baseTextureSampler;
	uniform sampler2D lightmapTextureSampler;


	in vec2 UV;
	in vec2 LightmapUV;

    out vec4 frag_colour;

	// Basetexture
	// Nothing is renderable without a base texture
	void GetBasetexture(inout vec4 fragColour, in sampler2D basetexture, in vec2 uv) 
	{
		fragColour = texture(basetexture, uv).rgba;
	}

	// Lightmaps the face
	// Does nothing if lightmap was not defined
	void ApplyLightmap(inout vec4 fragColour, in sampler2D lightmap, in vec2 uv) 
	{
		if (useLightmap == 0) {
			return;
		}

		fragColour = fragColour * texture(lightmap, uv).rgba;
	}

    void main() 
	{
		GetBasetexture(frag_colour, baseTextureSampler, UV);
		ApplyLightmap(frag_colour, lightmapTextureSampler, LightmapUV);
    }
` + "\x00"
