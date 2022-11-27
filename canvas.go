package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Canvas struct {
	Rect      sdl.Rect
	InsetRect sdl.Rect
	PixelSize int32

	ActivePixelIndex    int32
	ActivePixelPosition sdl.Point
}

func NewCanvas() (result Canvas) {
	result = Canvas{
		Rect:      sdl.Rect{X: 0, Y: 84, W: 404, H: 404},
		InsetRect: sdl.Rect{X: 10, Y: 94, W: 384, H: 384},
		PixelSize: 16,

		ActivePixelIndex: -1,
	}

	return
}

func (canvas *Canvas) Tick(input *Input, app *App) {
	if input.MousePosition.InRect(&canvas.InsetRect) {
		x := (input.MousePosition.X - canvas.InsetRect.X) / canvas.PixelSize
		y := (input.MousePosition.Y - canvas.InsetRect.Y) / canvas.PixelSize

		canvas.ActivePixelIndex = y*GRID_SIZE + x
		canvas.ActivePixelPosition = sdl.Point{X: canvas.InsetRect.X + x*canvas.PixelSize, Y: canvas.InsetRect.Y + y*canvas.PixelSize}

		if input.LMB == BS_JUST_PRESSED {
			app.StartInteraction(uint16(canvas.ActivePixelIndex))
		} else if input.LMB == BS_PRESSED {
			app.Interact(uint16(canvas.ActivePixelIndex))
		}
	} else {
		canvas.ActivePixelIndex = -1
	}
}

func (canvas *Canvas) Render(renderer *sdl.Renderer, app *App) {
	DrawRectOutset(renderer, Rect3DRenderData{
		Rect:           canvas.Rect,
		Color:          app.Theme["main"],
		ShadowColor:    app.Theme["shadow"],
		HighlightColor: app.Theme["highlight"],
	})

	DrawRectInset(renderer, Rect3DRenderData{
		Rect:           ExpandRect(canvas.InsetRect, 1),
		Color:          sdl.Color{R: 207, G: 207, B: 207, A: 207},
		ShadowColor:    app.Theme["shadow"],
		HighlightColor: app.Theme["highlight"],
	})

	for y := 0; y < int(GRID_SIZE); y++ {
		for x := 0; x < int(GRID_SIZE); x++ {
			if (x+y)%2 == 1 {
				DrawRect(renderer, RectRenderData{
					Rect:  sdl.Rect{X: canvas.InsetRect.X + int32(x)*canvas.PixelSize, Y: canvas.InsetRect.Y + int32(y)*canvas.PixelSize, W: canvas.PixelSize, H: canvas.PixelSize},
					Color: sdl.Color{R: 174, G: 171, B: 171, A: 255},
				})
			}
		}
	}

	activeFrame := app.Project.Frames[app.ActiveFrame]
	for y := 0; y < int(GRID_SIZE); y++ {
		for x := 0; x < int(GRID_SIZE); x++ {
			pixel := activeFrame[y*int(GRID_SIZE)+x]
			if pixel.A != 0 {
				DrawRect(renderer, RectRenderData{
					Rect:  sdl.Rect{X: canvas.InsetRect.X + int32(x)*canvas.PixelSize, Y: canvas.InsetRect.Y + int32(y)*canvas.PixelSize, W: canvas.PixelSize, H: canvas.PixelSize},
					Color: pixel,
				})
			}
		}
	}

	for _, idx := range app.Project.Selection {
		x := int32(idx) % GRID_SIZE
		y := int32(idx) / GRID_SIZE
		DrawRectOutline(renderer, RectOutlineRenderData{
			Rect:  sdl.Rect{X: canvas.InsetRect.X + x*canvas.PixelSize, Y: canvas.InsetRect.Y + y*canvas.PixelSize, W: canvas.PixelSize, H: canvas.PixelSize},
			Color: sdl.Color{R: 255, G: 255, B: 255, A: 255},
			Width: 1,
		})
	}

	if canvas.ActivePixelIndex >= 0 {
		DrawRectOutline(renderer, RectOutlineRenderData{
			Rect:  sdl.Rect{X: canvas.ActivePixelPosition.X, Y: canvas.ActivePixelPosition.Y, W: canvas.PixelSize, H: canvas.PixelSize},
			Color: sdl.Color{R: 255, G: 255, B: 255, A: 255},
			Width: 1,
		})
	}
}
