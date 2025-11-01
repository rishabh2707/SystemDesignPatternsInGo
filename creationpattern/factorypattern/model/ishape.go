package model

type IShape interface {
	Area() float64
	Perimeter() float64
	ToString() string
}
