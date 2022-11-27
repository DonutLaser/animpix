package main

import "github.com/veandco/go-sdl2/sdl"

type MenuItemType uint8

const (
	MIT_FILE MenuItemType = iota
	MIT_FRAME
	MIT_PALETTE
	MIT_NONE
)

type Menu struct {
	Rect sdl.Rect

	Buttons    []*MenuButton
	ActiveMenu MenuItemType
}

func NewMenu(app *App) (result Menu) {
	result = Menu{
		Rect: sdl.Rect{X: 0, Y: 0, W: 524, H: 32},

		Buttons:    make([]*MenuButton, 3),
		ActiveMenu: MIT_NONE,
	}

	font := app.Fonts["14px"]

	var startX int32 = result.Rect.X + 5
	result.Buttons[MIT_FILE] = NewMenuButton("File", &font, &sdl.Point{X: startX, Y: 6})
	result.Buttons[MIT_FILE].AddMenuItem("New Animation", &font)
	result.Buttons[MIT_FILE].AddMenuItem("Open Animation...", &font)
	result.Buttons[MIT_FILE].AddMenuItem("Save Animation", &font)
	result.Buttons[MIT_FILE].AddMenuItem("Save Animation as...", &font)

	startX += result.Buttons[MIT_FILE].Rect.W
	result.Buttons[MIT_FRAME] = NewMenuButton("Frame", &font, &sdl.Point{X: startX, Y: 6})
	result.Buttons[MIT_FRAME].AddMenuItem("Duplicate Frame", &font)
	result.Buttons[MIT_FRAME].AddMenuItem("Clear Frame", &font)

	startX += result.Buttons[MIT_FRAME].Rect.W
	result.Buttons[MIT_PALETTE] = NewMenuButton("Palette", &font, &sdl.Point{X: startX, Y: 6})
	result.Buttons[MIT_PALETTE].AddMenuItem("Default", &font)

	return
}

func (menu *Menu) Tick(input *Input, app *App) {
	for i := MIT_FILE; i <= MIT_PALETTE; i++ {
		if menu.Buttons[i].Tick(input) {
			if menu.ActiveMenu != MIT_NONE {
				menu.Buttons[menu.ActiveMenu].SetActive(false)
			}

			menu.ActiveMenu = i
		}
	}
}

func (menu *Menu) Render(renderer *sdl.Renderer, app *App) {
	DrawRectOutset(renderer, Rect3DRenderData{
		Rect:           menu.Rect,
		Color:          app.Theme["main"],
		ShadowColor:    app.Theme["shadow"],
		HighlightColor: app.Theme["highlight"],
	})

	for _, btn := range menu.Buttons {
		btn.Render(renderer, app)
	}
}
