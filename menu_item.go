package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type MenuItem struct {
	Rect      sdl.Rect
	Font      *Font
	TextWidth int32

	Text string
}

func NewMenuItem(text string, font *Font, position *sdl.Point) *MenuItem {
	return &MenuItem{
		Rect:      sdl.Rect{X: position.X, Y: position.Y, W: 128, H: 20},
		Font:      font,
		TextWidth: font.GetStringWidth(text),
		Text:      text,
	}
}

func (item *MenuItem) SetWidth(width int32) {
	item.Rect.W = width
}

func (item *MenuItem) Tick(input *Input, app *App) {
}

func (item *MenuItem) Render(renderer *sdl.Renderer, app *App) {
	DrawRect(renderer, RectRenderData{
		Rect:  item.Rect,
		Color: app.Theme["main"],
	})

	DrawText(renderer, TextRenderData{
		Rect: sdl.Rect{
			X: item.Rect.X + 5,
			Y: item.Rect.Y + (int32(item.Rect.H) / 2) - item.Font.Size/2,
			W: item.TextWidth,
			H: item.Font.Size,
		},
		Color: app.Theme["icon"],
		Font:  item.Font,
		Text:  item.Text,
	})
}
