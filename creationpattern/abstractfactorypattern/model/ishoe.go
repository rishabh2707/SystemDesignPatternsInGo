package model

import "fmt"

type IShoe interface {
	GetSize() int
	GetPrice() float64
	String() string
}

type Shoe struct {
	Size  int
	Price float64
}

func (s *Shoe) GetSize() int {
	return s.Size
}

func (s *Shoe) GetPrice() float64 {
	return s.Price
}

func (s *Shoe) SetSize(size int) {
	s.Size = size
}

func (s *Shoe) SetPrice(price float64) {
	s.Price = price
}

func (s *Shoe) String() string {
	return fmt.Sprintf("Shoe{Size: %d, Price: %f}", s.Size, s.Price)
}
