package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Colors struct {
	Rect            sdl.Rect
	PaletteRect     sdl.Rect
	ActiveColorRect sdl.Rect
	ColorsPerRow    int
	SwatchSize      int32

	Palette          []sdl.Color
	ActiveColorIndex uint16
}

func NewColors() (result Colors) {
	result = Colors{
		Rect:            sdl.Rect{X: 404, Y: 84, W: 120, H: 404},
		PaletteRect:     sdl.Rect{X: 414, Y: 94, W: 100, H: 342},
		ActiveColorRect: sdl.Rect{X: 414, Y: 446, W: 100, H: 32},
		ColorsPerRow:    5,
		SwatchSize:      20,

		Palette:          []sdl.Color{},
		ActiveColorIndex: 0,
	}

	return
}

func (colors *Colors) GetActiveColor() *sdl.Color {
	return &colors.Palette[colors.ActiveColorIndex]
}

func (colors *Colors) Tick(input *Input) {
	if input.MousePosition.InRect(&colors.PaletteRect) {
		if input.LMB == BS_JUST_PRESSED {
			x := (input.MousePosition.X - colors.PaletteRect.X) / colors.SwatchSize
			y := (input.MousePosition.Y - colors.PaletteRect.Y) / colors.SwatchSize

			newIndex := y*int32(colors.ColorsPerRow) + x
			if int(newIndex) < len(colors.Palette) {
				colors.ActiveColorIndex = uint16(newIndex)
			}
		}
	}
}

func (colors *Colors) Render(renderer *sdl.Renderer, app *App) {
	DrawRectOutset(renderer, Rect3DRenderData{
		Rect:           colors.Rect,
		Color:          app.Theme["main"],
		ShadowColor:    app.Theme["shadow"],
		HighlightColor: app.Theme["highlight"],
	})

	DrawRectInset(renderer, Rect3DRenderData{
		Rect:           ExpandRect(colors.PaletteRect, 1),
		Color:          app.Theme["inset"],
		ShadowColor:    app.Theme["shadow"],
		HighlightColor: app.Theme["highlight"],
	})

	DrawRectInset(renderer, Rect3DRenderData{
		Rect:           ExpandRect(colors.ActiveColorRect, 1),
		Color:          colors.Palette[colors.ActiveColorIndex],
		ShadowColor:    app.Theme["shadow"],
		HighlightColor: app.Theme["highlight"],
	})

	startX := colors.PaletteRect.X
	startY := colors.PaletteRect.Y
	index := 0
	for index < len(colors.Palette) {
		for i := 0; i < colors.ColorsPerRow; i++ {
			DrawRect(renderer, RectRenderData{
				Rect:  sdl.Rect{X: startX + int32(i)*colors.SwatchSize, Y: startY, W: colors.SwatchSize, H: colors.SwatchSize},
				Color: colors.Palette[index],
			})

			if index == int(colors.ActiveColorIndex) {
				DrawRectOutline(renderer, RectOutlineRenderData{
					Rect:  sdl.Rect{X: startX + int32(i)*colors.SwatchSize, Y: startY, W: colors.SwatchSize, H: colors.SwatchSize},
					Color: sdl.Color{R: 255, G: 255, B: 255, A: 255},
					Width: 2,
				})
			}

			index++
			if index == len(colors.Palette) {
				break
			}
		}

		startY += colors.SwatchSize
	}
}

func (colors *Colors) SetPalette(palette []sdl.Color) {
	colors.Palette = palette
}
