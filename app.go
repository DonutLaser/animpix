package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type ControlIconType uint8
type ToolType uint8

const (
	CIT_START_FRAME ControlIconType = iota
	CIT_PREV_FRAME
	CIT_PLAY
	CIT_NEXT_FRAME
	CIT_END_FRAME
	CIT_NEW_FRAME
)

const (
	TT_BRUSH ToolType = iota
	TT_ERASER
	TT_BUCKET
	TT_MOVE
	TT_SELECT
)

type App struct {
	Menu     Menu
	Controls Controls
	Canvas   Canvas
	Colors   Colors
	Toolbar  Toolbar

	ActiveTool  ToolType
	Project     Project
	ActiveFrame uint16

	ControlImages []Image
	ToolImages    []Image
	Theme         map[string]sdl.Color
}

var GRID_SIZE int32 = 24

func NewApp(windowWidth int32, windowHeight int32, renderer *sdl.Renderer) (result App) {
	result = App{
		Menu:   NewMenu(),
		Canvas: NewCanvas(),
		Colors: NewColors(),

		ActiveTool:  TT_BRUSH,
		Project:     NewProject(),
		ActiveFrame: 0,

		ControlImages: make([]Image, 6),
		ToolImages:    make([]Image, 5),
		Theme:         make(map[string]sdl.Color),
	}

	result.ControlImages[CIT_START_FRAME] = LoadImage("assets/icons/start_frame.png", renderer)
	result.ControlImages[CIT_PREV_FRAME] = LoadImage("assets/icons/prev_frame.png", renderer)
	result.ControlImages[CIT_PLAY] = LoadImage("assets/icons/play.png", renderer)
	result.ControlImages[CIT_NEXT_FRAME] = LoadImage("assets/icons/next_frame.png", renderer)
	result.ControlImages[CIT_END_FRAME] = LoadImage("assets/icons/end_frame.png", renderer)
	result.ControlImages[CIT_NEW_FRAME] = LoadImage("assets/icons/add_frame.png", renderer)

	result.ToolImages[TT_BRUSH] = LoadImage("assets/icons/brush.png", renderer)
	result.ToolImages[TT_ERASER] = LoadImage("assets/icons/eraser.png", renderer)
	result.ToolImages[TT_BUCKET] = LoadImage("assets/icons/bucket.png", renderer)
	result.ToolImages[TT_MOVE] = LoadImage("assets/icons/move.png", renderer)
	result.ToolImages[TT_SELECT] = LoadImage("assets/icons/select.png", renderer)

	result.Theme["main"] = sdl.Color{R: 47, G: 52, B: 61, A: 255}
	result.Theme["main_light"] = sdl.Color{R: 56, G: 64, B: 75, A: 255}
	result.Theme["inset"] = sdl.Color{R: 28, G: 33, B: 46, A: 255}
	result.Theme["shadow"] = sdl.Color{R: 23, G: 26, B: 30, A: 255}
	result.Theme["highlight"] = sdl.Color{R: 68, G: 73, B: 81, A: 255}
	result.Theme["icon"] = sdl.Color{R: 120, G: 131, B: 157, A: 255}

	result.Controls = NewControls(&result)
	result.Toolbar = NewToolbar(&result)
	result.Colors.SetPalette([]sdl.Color{
		{R: 0, G: 0, B: 0, A: 255},
		{R: 255, G: 0, B: 0, A: 255},
		{R: 0, G: 255, B: 0, A: 255},
		{R: 0, G: 0, B: 255, A: 255},
		{R: 255, G: 255, B: 0, A: 255},
		{R: 255, G: 0, B: 255, A: 255},
		{R: 0, G: 255, B: 255, A: 255},
		{R: 255, G: 255, B: 255, A: 255},
	})

	return
}

func (app *App) SelectTool(t ToolType) {
	app.ActiveTool = t
}

func (app *App) StartInteraction(pixelIndex uint16) {
	app.Project.ClearSelection()
	if app.ActiveTool == TT_BUCKET {
		app.Project.FillPixels(app.ActiveFrame, pixelIndex, app.Colors.GetActiveColor())
	} else {
		app.Interact(pixelIndex)
	}
}

func (app *App) Interact(pixelIndex uint16) {
	if app.ActiveTool == TT_BRUSH {
		app.Project.PaintPixel(app.ActiveFrame, pixelIndex, app.Colors.GetActiveColor())
	} else if app.ActiveTool == TT_ERASER {
		app.Project.ErasePixel(app.ActiveFrame, pixelIndex)
	} else if app.ActiveTool == TT_SELECT {
		app.Project.SelectPixel(pixelIndex)
	}
}

func (app *App) GoToFirstFrame() {
	app.ActiveFrame = 0
}

func (app *App) GoToPrevFrame() {
	if app.ActiveFrame > 0 {
		app.ActiveFrame -= 1
	}
}

func (app *App) Play() {
	fmt.Println("Playing")
}

func (app *App) GoToNextFrame() {
	if int(app.ActiveFrame) < len(app.Project.Frames)-1 {
		app.ActiveFrame += 1
	}
}

func (app *App) GoToLastFrame() {
	app.ActiveFrame = uint16(len(app.Project.Frames) - 1)
}

func (app *App) CreateNewFrame() {
	app.Project.AddNewFrame()
	app.GoToLastFrame()
}

func (app *App) Tick(input *Input) {
	if input.B == BS_JUST_PRESSED {
		app.Toolbar.SetButtonActive(TT_BRUSH, app.ActiveTool)
		app.SelectTool(TT_BRUSH)
	} else if input.E == BS_JUST_PRESSED {
		app.Toolbar.SetButtonActive(TT_ERASER, app.ActiveTool)
		app.SelectTool(TT_ERASER)
	} else if input.M == BS_JUST_PRESSED {
		app.Toolbar.SetButtonActive(TT_SELECT, app.ActiveTool)
		app.SelectTool(TT_SELECT)
	} else if input.V == BS_JUST_PRESSED {
		app.Toolbar.SetButtonActive(TT_MOVE, app.ActiveTool)
		app.SelectTool(TT_MOVE)
	} else if input.G == BS_JUST_PRESSED {
		app.Toolbar.SetButtonActive(TT_BUCKET, app.ActiveTool)
		app.SelectTool(TT_BUCKET)
	}

	app.Controls.Tick(input, app)
	app.Canvas.Tick(input, app)
	app.Colors.Tick(input)
	app.Toolbar.Tick(input, app)
}

func (app *App) Render(renderer *sdl.Renderer) {
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()

	app.Menu.Render(renderer, app)
	app.Controls.Render(renderer, app)
	app.Canvas.Render(renderer, app)
	app.Colors.Render(renderer, app)
	app.Toolbar.Render(renderer, app)

	renderer.Present()
}

func (app *App) Close() {
	for _, img := range app.ControlImages {
		img.Unload()
	}

	for _, img := range app.ToolImages {
		img.Unload()
	}
}
