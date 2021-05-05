package myopengl

import (
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func InitGlfw(width, height int) *glfw.Window {
	err := glfw.Init()
	if err != nil {
		log.Fatal("Error initializing glfw!")
	}

	window, err := glfw.CreateWindow(width, height, "something", nil, nil)
	if err != nil {
		log.Fatal("Error creating window")
	}

	window.MakeContextCurrent()
	err = gl.Init()
	if err != nil {
		log.Fatal("Error initializing opengl")
	}
	return window
}
