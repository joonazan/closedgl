package closedgl

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"log"
	"runtime"
)

func Run(render func(dt float64), width, height int, title string) {

	// OpenGL kaatuu jos sitä kutsutaan eri CPUista
	runtime.LockOSThread()

	complain(glfw.Init(), "Initializing GLFW:")
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.Resizable, glfw.False)
	w, err := glfw.CreateWindow(width, height, title, nil, nil)
	complain(err, "Creating window:")
	defer w.Destroy()

	w.MakeContextCurrent()
	complain(gl.Init(), "Initializing OpenGL:")

	for !w.ShouldClose() {

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		render(glfw.GetTime())

		w.SwapBuffers()
		glfw.PollEvents()
	}
}

func complain(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
	}
}
