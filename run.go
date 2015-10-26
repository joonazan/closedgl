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

	complain(glfw.Init())
	defer glfw.Terminate()

	complain(gl.Init())

	glfw.WindowHint(glfw.Resizable, glfw.False)
	w, err := glfw.CreateWindow(width, height, title, nil, nil)
	complain(err)
	defer w.Destroy()

	w.MakeContextCurrent()

	for !w.ShouldClose() {
		render(0)
		w.SwapBuffers()
		glfw.PollEvents()
	}
}

func complain(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
