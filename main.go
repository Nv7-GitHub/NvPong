package main

import (
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
)

func main() {
	runtime.LockOSThread()

	window := initGlfw()
	defer glfw.Terminate()
	program := initOpenGL()

	for !window.ShouldClose() {
		draw(window, program)
	}
}
