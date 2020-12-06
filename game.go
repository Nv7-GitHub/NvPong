package main

import (
	"math"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

var p1pos float32 = 0
var p2pos float32 = 0

var xvel float32 = 0.01
var yvel float32 = 0.02
var xpos float32 = 0
var ypos float32 = 0

func draw(window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	gl.Begin(gl.QUADS)

	// P1
	gl.Vertex2f(-0.975, p1pos)
	gl.Vertex2f(-0.975, p1pos+0.25)
	gl.Vertex2f(-0.95, p1pos+0.25)
	gl.Vertex2f(-0.95, p1pos)

	// P2
	gl.Vertex2f(0.975, p2pos)
	gl.Vertex2f(0.975, p2pos+0.25)
	gl.Vertex2f(0.95, p2pos+0.25)
	gl.Vertex2f(0.95, p2pos)

	gl.End()

	// Ball
	xpos += xvel
	ypos += yvel
	if (xpos >= 0.94) || (xpos <= -0.94) {
		xvel *= -1
	}
	if (ypos >= 0.875) || (ypos <= -0.875) {
		yvel *= -1
	}
	gl.Begin(gl.TRIANGLE_FAN)

	var triAmount float64 = 50
	gl.Vertex2f(xpos, ypos)
	for i := 0; float64(i) <= triAmount; i++ {
		gl.Vertex2f(
			xpos+(float32((math.Cos(float64(i)*math.Pi*2/triAmount)))/20),
			ypos+(float32((math.Sin(float64(i)*math.Pi*2/triAmount)))/10),
		)
	}

	gl.End()

	// Input
	if isPressed(window, glfw.KeyW) {
		p1pos += 0.05
	} else if isPressed(window, glfw.KeyS) {
		p1pos -= 0.05
	}

	if isPressed(window, glfw.KeyUp) {
		p2pos += 0.05
	} else if isPressed(window, glfw.KeyDown) {
		p2pos -= 0.05
	}

	glfw.PollEvents()
	window.SwapBuffers()
}
