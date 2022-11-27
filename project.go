package main

import "github.com/veandco/go-sdl2/sdl"

type Project struct {
	Frames    [][]sdl.Color
	Selection []uint16
}

type PixelToFill struct {
	x1 int32
	x2 int32
	y  int32
	dy int32
}

func NewProject() (result Project) {
	result = Project{
		Frames:    make([][]sdl.Color, 1),
		Selection: make([]uint16, 0),
	}

	result.Frames[0] = make([]sdl.Color, GRID_SIZE*GRID_SIZE)
	for i := 0; i < int(GRID_SIZE*GRID_SIZE); i++ {
		result.Frames[0][i] = sdl.Color{}
	}

	return
}

func (proj *Project) PaintPixel(frame uint16, pixelIndex uint16, color *sdl.Color) {
	proj.Frames[frame][pixelIndex].R = color.R
	proj.Frames[frame][pixelIndex].G = color.G
	proj.Frames[frame][pixelIndex].B = color.B
	proj.Frames[frame][pixelIndex].A = color.A
}

func (proj *Project) ErasePixel(frame uint16, pixelIndex uint16) {
	proj.Frames[frame][pixelIndex].R = 0
	proj.Frames[frame][pixelIndex].G = 0
	proj.Frames[frame][pixelIndex].B = 0
	proj.Frames[frame][pixelIndex].A = 0
}

// https://en.wikipedia.org/wiki/Flood_fill
func (proj *Project) FillPixels(frame uint16, pixelIndex uint16, color *sdl.Color) {
	colorToReplace := proj.Frames[frame][pixelIndex]

	if AreColorsEqual(&colorToReplace, color) {
		return
	}

	startX, startY := indexToPosition(pixelIndex, GRID_SIZE)
	stack := []PixelToFill{
		{x1: startX, x2: startX, y: startY, dy: 1},
		{x1: startX, x2: startX, y: startY - 1, dy: -1},
	}

	var pixel PixelToFill
	for len(stack) > 0 {
		pixel, stack = stack[len(stack)-1], stack[:len(stack)-1]
		x := pixel.x1
		if proj.isInside(x, pixel.y, frame, &colorToReplace) {
			for proj.isInside(x-1, pixel.y, frame, &colorToReplace) {
				proj.PaintPixel(frame, positionToIndex(x-1, pixel.y, GRID_SIZE), color)
				x -= 1
			}
		}

		if x < pixel.x1 {
			stack = append(stack, PixelToFill{x1: x, x2: pixel.x1 - 1, y: pixel.y - pixel.dy, dy: -pixel.dy})
		}

		for pixel.x1 <= pixel.x2 {
			for proj.isInside(pixel.x1, pixel.y, frame, &colorToReplace) {
				proj.PaintPixel(frame, positionToIndex(pixel.x1, pixel.y, GRID_SIZE), color)
				pixel.x1 += 1
				stack = append(stack, PixelToFill{x1: x, x2: pixel.x1 - 1, y: pixel.y + pixel.dy, dy: pixel.dy})
				if pixel.x1-1 > pixel.x2 {
					stack = append(stack, PixelToFill{x1: pixel.x2 + 1, x2: pixel.x1 - 1, y: pixel.y - pixel.dy, dy: -pixel.dy})
				}
			}

			pixel.x1 += 1
			for pixel.x1 < pixel.x2 && !proj.isInside(pixel.x1, pixel.y, frame, &colorToReplace) {
				pixel.x1 += 1
			}

			x = pixel.x1
		}
	}
}

func (proj *Project) ClearSelection() {
	proj.Selection = nil
}

func (proj *Project) SelectPixel(pixelIndex uint16) {
	proj.Selection = append(proj.Selection, pixelIndex)
}

func (proj *Project) AddNewFrame() {
	proj.ClearSelection()

	proj.Frames = append(proj.Frames, []sdl.Color{})
	proj.Frames[len(proj.Frames)-1] = make([]sdl.Color, GRID_SIZE*GRID_SIZE)
	for i := 0; i < int(GRID_SIZE*GRID_SIZE); i++ {
		proj.Frames[len(proj.Frames)-1][i] = sdl.Color{}
	}
}

func (proj *Project) isInside(x int32, y int32, frame uint16, color *sdl.Color) bool {
	return x >= 0 && x < GRID_SIZE && y >= 0 && y < GRID_SIZE && AreColorsEqual(&proj.Frames[frame][positionToIndex(x, y, GRID_SIZE)], color)
}
