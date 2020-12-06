package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

var x float32 = 0
var y float32 = 0

func draw(window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	gl.Begin(gl.QUADS)
	gl.Vertex2f(x, y)
	gl.Vertex2f(x, y+0.5)
	gl.Vertex2f(x+0.5, y+0.5)
	gl.Vertex2f(x+0.5, y)
	gl.End()

	if isPressed(window, glfw.KeyRight) {
		x += 0.01
	} else if isPressed(window, glfw.KeyLeft) {
		x -= 0.01
	}

	if isPressed(window, glfw.KeyUp) {
		y += 0.01
	} else if isPressed(window, glfw.KeyDown) {
		y -= 0.01
	}

	glfw.PollEvents()
	window.SwapBuffers()
}
