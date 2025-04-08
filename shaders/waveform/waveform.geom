#version 330 core

layout(lines) in; // Input primitive type (expects, for instance, lines)
layout(triangle_strip, max_vertices = 4) out; // Output type and max vertices emitted

uniform float thickness; // Line thickness uniform

void main() {
    // Calculate the direction vector for the line
    vec2 direction = normalize(gl_in[1].gl_Position.xy - gl_in[0].gl_Position.xy);
    vec2 perp = vec2(-direction.y, direction.x) * thickness * 0.5;  // Perpendicular vector

    // Emit a triangle strip for the thick line
    gl_Position = gl_in[0].gl_Position + vec4(perp, 0.0, 0.0);
    EmitVertex();

    gl_Position = gl_in[0].gl_Position - vec4(perp, 0.0, 0.0);
    EmitVertex();

    gl_Position = gl_in[1].gl_Position + vec4(perp, 0.0, 0.0);
    EmitVertex();

    gl_Position = gl_in[1].gl_Position - vec4(perp, 0.0, 0.0);
    EmitVertex();

    EndPrimitive();
}
