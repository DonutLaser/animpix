package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Controls struct {
	Rect         sdl.Rect
	Spacing      int32
	ButtonWidth  int32
	ButtonHeight int32

	Buttons []*Button
}

func NewControls(app *App) (result Controls) {
	result = Controls{
		Rect:         sdl.Rect{X: 0, Y: 32, W: 524, H: 52},
		Spacing:      16,
		ButtonWidth:  64,
		ButtonHeight: 32,

		Buttons: make([]*Button, 6),
	}

	var btnCount int32 = int32(len(result.Buttons))

	var startX int32 = (result.Rect.W - (result.ButtonWidth*btnCount + result.Spacing*(btnCount-1))) / 2
	for i := CIT_START_FRAME; i <= CIT_NEW_FRAME; i++ {
		result.Buttons[i] = NewButton(&app.ControlImages[i], &sdl.Point{X: startX + 1, Y: result.Rect.Y + 11})
		startX += result.ButtonWidth + result.Spacing
	}

	return
}

func (controls *Controls) Tick(input *Input, app *App) {
	if controls.Buttons[CIT_START_FRAME].Tick(input) {
		app.GoToFirstFrame()
	} else if controls.Buttons[CIT_PREV_FRAME].Tick(input) {
		app.GoToPrevFrame()
	} else if controls.Buttons[CIT_PLAY].Tick(input) {
		app.Play()
	} else if controls.Buttons[CIT_NEXT_FRAME].Tick(input) {
		app.GoToNextFrame()
	} else if controls.Buttons[CIT_END_FRAME].Tick(input) {
		app.GoToLastFrame()
	} else if controls.Buttons[CIT_NEW_FRAME].Tick(input) {
		app.CreateNewFrame()
	}
}

func (controls *Controls) Render(renderer *sdl.Renderer, app *App) {
	DrawRectOutset(renderer, Rect3DRenderData{
		Rect:           controls.Rect,
		Color:          app.Theme["main"],
		ShadowColor:    app.Theme["shadow"],
		HighlightColor: app.Theme["highlight"],
	})

	var startX int32 = (controls.Rect.W - (controls.ButtonWidth*6 + controls.Spacing*5)) / 2
	for i := 0; i < len(controls.Buttons); i++ {
		DrawRectInset(renderer, Rect3DRenderData{
			Rect:           sdl.Rect{X: startX, Y: controls.Rect.Y + (controls.Rect.H / 2) - controls.ButtonHeight/2, W: 66, H: controls.ButtonHeight},
			Color:          app.Theme["inset"],
			ShadowColor:    app.Theme["shadow"],
			HighlightColor: app.Theme["highlight"],
		})

		controls.Buttons[i].Render(renderer, app)

		startX += int32(controls.ButtonWidth + controls.Spacing)
	}
}
