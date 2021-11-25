package models

import "lesson2/colours"

type Point struct {
	Y, X int
	Colour colours.Colour
}

func (p *Point) SetColour(colour colours.Colour)  {
	p.Colour = colour
}