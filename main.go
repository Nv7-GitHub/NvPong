package main

import (
	"runtime"

	// OR: github.com/go-gl/gl/v2.1/gl
	"github.com/go-gl/glfw/v3.2/glfw"
)

func main() {
	runtime.LockOSThread()

	window := initGlfw()
	defer glfw.Terminate()
	program := initOpenGL()

	vao := makeVao(shape)
	for !window.ShouldClose() {
		draw(vao, window, program)
	}
}
