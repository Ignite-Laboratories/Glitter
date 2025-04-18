#version 310 es

layout (location = 0) in vec3 aPos;

uniform mat4 uProjectionMatrix;

void main() {
    gl_Position = uProjectionMatrix * vec4(aPos, 1.0);
}
