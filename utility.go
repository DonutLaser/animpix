package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

func FailIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ExpandRect(rect sdl.Rect, by int32) sdl.Rect {
	return sdl.Rect{
		X: rect.X - by,
		Y: rect.Y - by,
		W: rect.W + by*2,
		H: rect.H + by*2,
	}
}

func AreColorsEqual(color1 *sdl.Color, color2 *sdl.Color) bool {
	return color1.R == color2.R && color1.G == color2.G && color1.B == color2.B && color1.A == color2.A
}

func indexToPosition(index uint16, gridSize int32) (int32, int32) {
	return int32(index) % gridSize, int32(index) / gridSize
}

func positionToIndex(x int32, y int32, gridSize int32) uint16 {
	return uint16(y*gridSize + x)
}
