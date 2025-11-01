package main

import (
	"fmt"

	"com.systemdesignpatterns.go.factorypattern/factory"
)

func main() {
	rectangle := factory.GetShape("rectangle")
	fmt.Println(rectangle.ToString())
	fmt.Println("Area:", rectangle.Area())
	fmt.Println("Perimeter:", rectangle.Perimeter())
	circle := factory.GetShape("circle")
	fmt.Println(circle.ToString())
	fmt.Println("Area:", circle.Area())
	fmt.Println("Perimeter:", circle.Perimeter())
	square := factory.GetShape("square")
	fmt.Println(square.ToString())
	fmt.Println("Area:", square.Area())
	fmt.Println("Perimeter:", square.Perimeter())
}
