package main

import (
	"fmt"

	"com.systemdesignpatterns.go.abstractfactorypattern/factory"
)

func main() {
	adidasFactory := factory.GetFactory("Adidas")
	adidasShirt := adidasFactory.CreateShirt()
	adidasShoe := adidasFactory.CreateShoe()
	fmt.Println(adidasShirt.String())
	fmt.Println(adidasShoe.String())

	nikeFactory := factory.GetFactory("Nike")
	nikeShirt := nikeFactory.CreateShirt()
	nikeShoe := nikeFactory.CreateShoe()
	fmt.Println(nikeShirt.String())
	fmt.Println(nikeShoe.String())
}
