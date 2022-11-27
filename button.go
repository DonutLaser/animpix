package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Button struct {
	Rect     sdl.Rect
	Icon     *Image
	IsHot    bool
	IsActive bool
}

func NewButton(image *Image, position *sdl.Point) *Button {
	return &Button{
		Rect: sdl.Rect{X: position.X, Y: position.Y, W: 64, H: 30},
		Icon: image,
	}
}

func (button *Button) Tick(input *Input) bool {
	result := false

	if input.MousePosition.InRect(&button.Rect) {
		button.IsHot = true

		if input.LMB == BS_JUST_PRESSED {
			button.IsActive = true
		} else if input.LMB == BS_RELEASED && button.IsActive {
			button.IsActive = false
			result = true
		}
	} else {
		if button.IsActive {
			if input.LMB == BS_RELEASED || input.LMB == BS_NONE {
				button.IsActive = false
			}
		} else {
			button.IsHot = false
		}
	}

	return result
}

func (button *Button) Render(renderer *sdl.Renderer, app *App) {
	buttonColor := app.Theme["main"]
	if button.IsHot {
		buttonColor = app.Theme["main_light"]
	}

	if !button.IsActive {
		DrawRectOutset(renderer, Rect3DRenderData{
			Rect:           button.Rect,
			Color:          buttonColor,
			ShadowColor:    app.Theme["shadow"],
			HighlightColor: app.Theme["highlight"],
		})
	}

	DrawImage(renderer, ImageRenderData{
		Rect: sdl.Rect{
			X: button.Rect.X + (int32(button.Rect.W) / 2) - button.Icon.Width/2,
			Y: button.Rect.Y + (int32(button.Rect.H) / 2) - button.Icon.Height/2,
			W: button.Icon.Width,
			H: button.Icon.Height,
		},
		Color: app.Theme["icon"],
		Image: button.Icon,
	})
}
