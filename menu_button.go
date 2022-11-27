package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type MenuButton struct {
	Rect      sdl.Rect
	Padding   int32
	NameWidth int32
	Font      *Font

	Name     string
	IsHot    bool
	IsActive bool
	IsOn     bool

	Items []*MenuItem
}

func NewMenuButton(name string, font *Font, position *sdl.Point) *MenuButton {
	nameWidth := font.GetStringWidth(name)
	var padding int32 = 7

	return &MenuButton{
		Rect:      sdl.Rect{X: position.X, Y: position.Y, W: nameWidth + padding*2, H: 20},
		Padding:   padding,
		NameWidth: nameWidth,
		Font:      font,

		Name: name,

		Items: make([]*MenuItem, 0),
	}
}

func (btn *MenuButton) SetActive(active bool) {
	btn.IsOn = active
}

func (btn *MenuButton) AddMenuItem(text string, font *Font) {
	btn.Items = append(btn.Items, NewMenuItem(text, font, &sdl.Point{
		X: btn.Rect.X,
		Y: btn.Rect.Y + btn.Rect.H + 2 + int32(len(btn.Items))*20,
	}))

	largestItemWidth := btn.Items[0].TextWidth
	for _, item := range btn.Items {
		if item.TextWidth > largestItemWidth {
			largestItemWidth = item.TextWidth
		}
	}

	for _, item := range btn.Items {
		item.SetWidth(largestItemWidth + 10)
	}
}

func (btn *MenuButton) Tick(input *Input) bool {
	result := false

	if input.MousePosition.InRect(&btn.Rect) {
		btn.IsHot = true

		if input.LMB == BS_JUST_PRESSED {
			btn.IsActive = true
		} else if input.LMB == BS_RELEASED && btn.IsActive {
			btn.IsActive = false

			if !btn.IsOn {
				result = true
				btn.IsOn = true
			}
		}
	} else {
		if btn.IsActive && input.LMB == BS_RELEASED {
			btn.IsActive = false
		}

		btn.IsHot = false
	}

	return result
}

func (btn *MenuButton) Render(renderer *sdl.Renderer, app *App) {
	if btn.IsActive || btn.IsOn {
		DrawRectInset(renderer, Rect3DRenderData{
			Rect:           btn.Rect,
			Color:          app.Theme["inset"],
			ShadowColor:    app.Theme["shadow"],
			HighlightColor: app.Theme["highlight"],
		})

		if btn.IsOn {
			for _, item := range btn.Items {
				item.Render(renderer, app)
			}

			DrawRectOutline(renderer, RectOutlineRenderData{
				Rect:  sdl.Rect{X: btn.Items[0].Rect.X, Y: btn.Items[0].Rect.Y, W: btn.Items[0].Rect.W, H: int32(len(btn.Items)) * 20},
				Color: app.Theme["inset"],
				Width: 1,
			})
		}
	} else if btn.IsHot {
		DrawRect(renderer, RectRenderData{
			Rect:  btn.Rect,
			Color: app.Theme["main_light"],
		})
	}

	DrawText(renderer, TextRenderData{
		Rect: sdl.Rect{
			X: btn.Rect.X + (int32(btn.Rect.W) / 2) - btn.NameWidth/2,
			Y: btn.Rect.Y + (int32(btn.Rect.H) / 2) - btn.Font.Size/2,
			W: btn.NameWidth,
			H: btn.Font.Size,
		},
		Color: app.Theme["icon"],
		Font:  btn.Font,
		Text:  btn.Name,
	})
}
