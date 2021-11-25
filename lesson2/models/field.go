package models

import "lesson2/colours"

type Field struct {
	Matrix [][]Point
}

func (f *Field) Init()  {
	for i := 0; i < len(f.Matrix); i++ {
		for j := 0; j < len(f.Matrix[0]); j++ {
			f.Matrix[i][j] = Point{Y: i, X: j, Colour: colours.WHITE}
		}
	}
}

func (f *Field) SetPoints(points [3]Point)  {
	for _, p := range points {
		f.Matrix[p.Y][p.X] = p
	}
}

