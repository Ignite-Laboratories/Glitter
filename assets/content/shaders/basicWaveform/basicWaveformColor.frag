#version 310 es
precision mediump float; // Specify floating-point precision

uniform vec4 fgColor;

out vec4 FragColor;

void main() {
    FragColor = fgColor;
}
