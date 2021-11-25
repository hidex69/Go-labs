package main

import (
	"fmt"
	"lesson2/colours"
	"lesson2/models"
)

func main() {
	var p = models.Point{X: 1, Y: 2, Colour: colours.WHITE}
	p.SetColour(colours.BLACK)
	fmt.Println(p)
}
