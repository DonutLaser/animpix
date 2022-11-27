package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Toolbar struct {
	Rect         sdl.Rect
	Height       uint16
	Spacing      int32
	ButtonWidth  int32
	ButtonHeight int32

	Buttons []*ToggleButton
}

func NewToolbar(app *App) (result Toolbar) {
	result = Toolbar{
		Rect:         sdl.Rect{X: 0, Y: 488, W: 524, H: 52},
		Spacing:      16,
		ButtonWidth:  32,
		ButtonHeight: 32,

		Buttons: make([]*ToggleButton, 5),
	}

	var btnCount int32 = int32(len(result.Buttons))

	var startX int32 = (result.Rect.W - (result.ButtonWidth*btnCount + result.Spacing*(btnCount-1))) / 2
	for i := TT_BRUSH; i <= TT_SELECT; i++ {
		result.Buttons[i] = NewToggleButton(&app.ToolImages[i], &sdl.Point{X: startX + 1, Y: result.Rect.Y + 11})
		startX += result.ButtonWidth + result.Spacing
	}

	result.Buttons[TT_BRUSH].SetActive(true)

	return
}

func (toolbar *Toolbar) SetButtonActive(btn ToolType, activeTool ToolType) {
	toolbar.Buttons[activeTool].SetActive(false)
	toolbar.Buttons[btn].SetActive(true)
}

func (toolbar *Toolbar) Tick(input *Input, app *App) {
	if toolbar.Buttons[TT_BRUSH].Tick(input) {
		toolbar.Buttons[app.ActiveTool].SetActive(false)
		app.SelectTool(TT_BRUSH)
	} else if toolbar.Buttons[TT_ERASER].Tick(input) {
		toolbar.Buttons[app.ActiveTool].SetActive(false)
		app.SelectTool(TT_ERASER)
	} else if toolbar.Buttons[TT_BUCKET].Tick(input) {
		toolbar.Buttons[app.ActiveTool].SetActive(false)
		app.SelectTool(TT_BUCKET)
	} else if toolbar.Buttons[TT_MOVE].Tick(input) {
		toolbar.Buttons[app.ActiveTool].SetActive(false)
		app.SelectTool(TT_MOVE)
	} else if toolbar.Buttons[TT_SELECT].Tick(input) {
		toolbar.Buttons[app.ActiveTool].SetActive(false)
		app.SelectTool(TT_SELECT)
	}
}

func (toolbar *Toolbar) Render(renderer *sdl.Renderer, app *App) {
	DrawRectOutset(renderer, Rect3DRenderData{
		Rect:           toolbar.Rect,
		Color:          app.Theme["main"],
		ShadowColor:    app.Theme["shadow"],
		HighlightColor: app.Theme["highlight"],
	})

	var startX int32 = (toolbar.Rect.W - (toolbar.ButtonWidth*5 + toolbar.Spacing*4)) / 2
	for i := 0; i < len(toolbar.Buttons); i++ {
		DrawRectInset(renderer, Rect3DRenderData{
			Rect:           sdl.Rect{X: startX, Y: toolbar.Rect.Y + (toolbar.Rect.H / 2) - toolbar.ButtonHeight/2, W: 34, H: toolbar.ButtonHeight},
			Color:          app.Theme["inset"],
			ShadowColor:    app.Theme["shadow"],
			HighlightColor: app.Theme["highlight"],
		})

		toolbar.Buttons[i].Render(renderer, app)

		startX += int32(toolbar.ButtonWidth + toolbar.Spacing)
	}
}
