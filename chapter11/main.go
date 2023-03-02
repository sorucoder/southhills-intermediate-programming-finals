package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

// This is a hacky way to do this. I only did so to satisfy the prompt.
// I would not normally do this.
func NewCircle(radius ...float64) *Circle {
	if len(radius) > 0 {
		return &Circle{radius[0]}
	} else {
		return &Circle{1}
	}
}

func (circle *Circle) Radius() float64 {
	return circle.radius
}

func (circle *Circle) SetRadius(radius float64) {
	circle.radius = radius
}

func (circle *Circle) Diameter() float64 {
	return 2.0 * circle.radius
}

func (circle *Circle) Area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func main() {
	manualCircle := NewCircle(2.5)
	defaultCircle := NewCircle()
	fmt.Printf("MANUAL:  radius = %2.4f, diameter = %2.4f, area = %2.4f\n", manualCircle.Radius(), manualCircle.Diameter(), manualCircle.Area())
	fmt.Printf("DEFAULT: radius = %2.4f, diameter = %2.4f, area = %2.4f\n", defaultCircle.Radius(), defaultCircle.Diameter(), defaultCircle.Area())
}
