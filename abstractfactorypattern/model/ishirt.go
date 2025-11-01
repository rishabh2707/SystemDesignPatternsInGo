package model

import "fmt"

type IShirt interface {
	GetColor() string
	GetSize() string
	GetPrice() float64
	String() string
}

type Shirt struct {
	Color string
	Size  string
	Price float64
}

func (s *Shirt) GetColor() string {
	return s.Color
}

func (s *Shirt) GetSize() string {
	return s.Size
}

func (s *Shirt) GetPrice() float64 {
	return s.Price
}

func (s *Shirt) SetColor(color string) {
	s.Color = color
}

func (s *Shirt) SetSize(size string) {
	s.Size = size
}

func (s *Shirt) SetPrice(price float64) {
	s.Price = price
}

func (s *Shirt) String() string {
	return fmt.Sprintf("Shirt{Color: %s, Size: %s, Price: %f}", s.Color, s.Size, s.Price)
}
