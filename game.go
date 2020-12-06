package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

var x = 0
var y = 0

func draw(window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	gl.Begin(gl.QUADS)
	gl.Vertex2f(0, 0)
	gl.Vertex2f(0, 0.5)
	gl.Vertex2f(0.5, 0.5)
	gl.Vertex2f(0.5, 0)
	gl.End()

	glfw.PollEvents()
	window.SwapBuffers()
}
