package main

import "github.com/veandco/go-sdl2/sdl"

type Menu struct {
	Rect sdl.Rect
}

func NewMenu() (result Menu) {
	result = Menu{
		Rect: sdl.Rect{X: 0, Y: 0, W: 524, H: 32},
	}

	return
}

func (menu *Menu) Tick(input *Input, app *App) {
}

func (menu *Menu) Render(renderer *sdl.Renderer, app *App) {
	DrawRectOutset(renderer, Rect3DRenderData{
		Rect:           menu.Rect,
		Color:          app.Theme["main"],
		ShadowColor:    app.Theme["shadow"],
		HighlightColor: app.Theme["highlight"],
	})
}
