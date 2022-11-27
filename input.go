package main

import "github.com/veandco/go-sdl2/sdl"

type ButtonState uint8

const (
	BS_NONE ButtonState = iota
	BS_JUST_PRESSED
	BS_PRESSED
	BS_RELEASED
)

type Input struct {
	MousePosition sdl.Point
	LMB           ButtonState
	RMB           ButtonState
	B             ButtonState
	E             ButtonState
	M             ButtonState
	V             ButtonState
	G             ButtonState
}

func (input *Input) Update() {
	input.LMB = input.updateButtonState(input.LMB)
	input.RMB = input.updateButtonState(input.RMB)
	input.B = input.updateButtonState(input.B)
	input.E = input.updateButtonState(input.E)
	input.M = input.updateButtonState(input.M)
	input.V = input.updateButtonState(input.V)
	input.G = input.updateButtonState(input.G)
}

func (input *Input) UpdateButton(btn ButtonState, newState uint8) ButtonState {
	if newState == sdl.RELEASED {
		return BS_RELEASED
	}

	if btn != BS_PRESSED {
		return BS_JUST_PRESSED
	}

	return btn
}

func (input *Input) updateButtonState(btn ButtonState) ButtonState {
	if btn == BS_RELEASED {
		btn = BS_NONE
	} else if btn == BS_JUST_PRESSED {
		btn = BS_PRESSED
	}

	return btn
}
