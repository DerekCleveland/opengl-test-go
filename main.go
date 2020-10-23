package main

import (
	"runtime"

	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	err = gl.Init()
	if err != nil {
		panic(err)
	}

	// Sets the background color to blue
	//gl.ClearColor(0, 0.5, 1.0, 1.0)

	for !window.ShouldClose() {
		// gl.Clear clears the screen
		//gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		myInit()
		myDisplay()
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func myInit() {
	gl.ClearColor(1.0, 1.0, 1.0, 0.0) // Background color to white
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0.0, 1.0, 0.0, 1.0, -1.0, 1.0)
	gl.
}

func myDisplay() {
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.LoadIdentity()
	gl.Rotatef(90.0, 0.0, 0.0, 1.0)
	gl.Translatef(0.0, -0.75, 0.0)
	drawDiode()

	gl.LoadIdentity()
	gl.Rotatef(90.0, 0.0, 0.0, 1.0)
	gl.Translatef(0.0, -1.25, 0.0)
	drawDiode()

	gl.LoadIdentity()
	gl.Translatef(0.0, -0.3, 0.0)
	drawDiode()

	gl.LoadIdentity()
	gl.Translatef(0.0, 0.3, 0.0)
	drawDiode()

	gl.LoadIdentity()
	//gl.ulookat //hmmmmm this doesnt exist
	gl.Flush()
	
}

func drawDiode() {
	gl.Color3f(1.0, 0.0, 0.0)
	gl.LineWidth(2.0)
	gl.Begin(gl.TRIANGLES)
	gl.Vertex3f(0.4, 0.5, 0.0)
	gl.Vertex3f(0.6, 0.6, 0.0)
	gl.Vertex3f(0.6, 0.4, 0.0)
	gl.End()

	gl.Color3f(0.0, 0.0, 0.0)
	gl.Begin(gl.LINES)
	gl.Vertex3f(0.4, 0.5, 0.0)
	gl.Vertex3f(0.6, 0.6, 0.0)

	gl.Vertex3f(0.6, 0.6, 0.0)
	gl.Vertex3f(0.6, 0.4, 0.0)

	gl.Vertex3f(0.4, 0.5, 0.0)
	gl.Vertex3f(0.6, 0.4, 0.0)

	gl.Vertex3f(0.4, 0.6, 0.0)
	gl.Vertex3f(0.4, 0.4, 0.0)

	gl.Vertex3f(0.25, 0.5, 0.0)
	gl.Vertex3f(0.4, 0.5, 0.0)

	gl.Vertex3f(0.6, 0.5, 0.0)
	gl.Vertex3f(0.75, 0.5, 0.0)
	gl.End()
}