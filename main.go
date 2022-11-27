package main

import (
	"fmt"
	_ "net/http/pprof"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	defer sdl.Quit()

	err = ttf.Init()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	defer ttf.Quit()

	window, err := sdl.CreateWindow("Animpix", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 524, 540, 0)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	defer renderer.Destroy()

	windowWidth, windowHeight := window.GetSize()

	app := NewApp(windowWidth, windowHeight, renderer)
	input := Input{}

	running := true
	for running {
		input.Update()

		event := sdl.WaitEvent()
		switch t := event.(type) {
		case *sdl.QuitEvent:
			running = false
		case *sdl.MouseMotionEvent:
			input.MousePosition.X = t.X
			input.MousePosition.Y = t.Y
		case *sdl.MouseButtonEvent:
			if t.Button == sdl.BUTTON_LEFT {
				input.LMB = input.UpdateButton(input.LMB, t.State)
			} else if t.Button == sdl.BUTTON_RIGHT {
				input.RMB = input.UpdateButton(input.RMB, t.State)
			}
		case *sdl.KeyboardEvent:
			keycode := t.Keysym.Sym

			switch keycode {
			case sdl.K_b:
				input.B = input.UpdateButton(input.B, t.State)
			case sdl.K_e:
				input.E = input.UpdateButton(input.E, t.State)
			case sdl.K_m:
				input.M = input.UpdateButton(input.M, t.State)
			case sdl.K_v:
				input.V = input.UpdateButton(input.V, t.State)
			case sdl.K_g:
				input.G = input.UpdateButton(input.G, t.State)
			}
		}

		app.Tick(&input)
		app.Render(renderer)
	}

	app.Close()
}
