package shaders


var Fragment = `
    #version 410
    out vec4 frag_colour;
    void main() {
        frag_colour = vec4(1, 1, 1, 0.025);
    }
` + "\x00"