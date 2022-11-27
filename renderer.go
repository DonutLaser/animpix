package main

import "github.com/veandco/go-sdl2/sdl"

type RectRenderData struct {
	Rect  sdl.Rect
	Color sdl.Color
}

type Rect3DRenderData struct {
	Rect           sdl.Rect
	Color          sdl.Color
	ShadowColor    sdl.Color
	HighlightColor sdl.Color
}

type RectOutlineRenderData struct {
	Rect  sdl.Rect
	Color sdl.Color
	Width uint8
}

type TextRenderData struct {
	Rect  sdl.Rect
	Color sdl.Color
	Font  *Font
	Text  string
}

type ImageRenderData struct {
	Rect  sdl.Rect
	Color sdl.Color
	Image *Image
}

func DrawRect(renderer *sdl.Renderer, data RectRenderData) {
	renderer.SetDrawColor(data.Color.R, data.Color.G, data.Color.B, data.Color.A)
	renderer.FillRect(&data.Rect)
}

func DrawRectOutset(renderer *sdl.Renderer, data Rect3DRenderData) {
	renderer.SetDrawColor(data.Color.R, data.Color.G, data.Color.B, data.Color.A)
	renderer.FillRect(&data.Rect)

	left := sdl.Rect{X: data.Rect.X, Y: data.Rect.Y, W: 1, H: data.Rect.H}
	top := sdl.Rect{X: data.Rect.X, Y: data.Rect.Y, W: data.Rect.W, H: 1}
	right := sdl.Rect{X: data.Rect.X + data.Rect.W - 1, Y: data.Rect.Y, W: 1, H: data.Rect.H}
	bottom := sdl.Rect{X: data.Rect.X, Y: data.Rect.Y + data.Rect.H - 1, W: data.Rect.W, H: 1}

	renderer.SetDrawColor(data.HighlightColor.R, data.HighlightColor.G, data.HighlightColor.B, data.HighlightColor.A)
	renderer.FillRect(&left)
	renderer.FillRect(&top)

	renderer.SetDrawColor(data.ShadowColor.R, data.ShadowColor.G, data.ShadowColor.B, data.ShadowColor.A)
	renderer.FillRect(&right)
	renderer.FillRect(&bottom)
}

func DrawRectInset(renderer *sdl.Renderer, data Rect3DRenderData) {
	renderer.SetDrawColor(data.Color.R, data.Color.G, data.Color.B, data.Color.A)
	renderer.FillRect(&data.Rect)

	left := sdl.Rect{X: data.Rect.X, Y: data.Rect.Y, W: 1, H: data.Rect.H}
	top := sdl.Rect{X: data.Rect.X, Y: data.Rect.Y, W: data.Rect.W, H: 1}
	right := sdl.Rect{X: data.Rect.X + data.Rect.W - 1, Y: data.Rect.Y, W: 1, H: data.Rect.H}
	bottom := sdl.Rect{X: data.Rect.X, Y: data.Rect.Y + data.Rect.H - 1, W: data.Rect.W, H: 1}

	renderer.SetDrawColor(data.ShadowColor.R, data.ShadowColor.G, data.ShadowColor.B, data.ShadowColor.A)
	renderer.FillRect(&left)
	renderer.FillRect(&top)

	renderer.SetDrawColor(data.HighlightColor.R, data.HighlightColor.G, data.HighlightColor.B, data.HighlightColor.A)
	renderer.FillRect(&right)
	renderer.FillRect(&bottom)
}

func DrawRectOutline(renderer *sdl.Renderer, data RectOutlineRenderData) {
	renderer.SetDrawColor(data.Color.R, data.Color.G, data.Color.B, data.Color.A)

	top := sdl.Rect{X: data.Rect.X, Y: data.Rect.Y, W: data.Rect.W, H: int32(data.Width)}
	right := sdl.Rect{X: data.Rect.X + data.Rect.W - int32(data.Width), Y: data.Rect.Y, W: int32(data.Width), H: data.Rect.H}
	bottom := sdl.Rect{X: data.Rect.X, Y: data.Rect.Y + data.Rect.H - int32(data.Width), W: data.Rect.W, H: int32(data.Width)}
	left := sdl.Rect{X: data.Rect.X, Y: data.Rect.Y, W: int32(data.Width), H: data.Rect.H}

	renderer.FillRect(&top)
	renderer.FillRect(&right)
	renderer.FillRect(&bottom)
	renderer.FillRect(&left)
}

func DrawText(renderer *sdl.Renderer, data TextRenderData) {
	surface, _ := data.Font.Data.RenderUTF8Blended(data.Text, data.Color)
	defer surface.Free()

	texture, _ := renderer.CreateTextureFromSurface(surface)
	defer texture.Destroy()

	renderer.Copy(texture, nil, &data.Rect)
}

func DrawImage(renderer *sdl.Renderer, data ImageRenderData) {
	data.Image.Data.SetColorMod(data.Color.R, data.Color.G, data.Color.B)
	renderer.Copy(data.Image.Data, nil, &data.Rect)
}
