package main

import "fmt"

func main() {
	margherita := NewMargheritaPizza()
	fmt.Println(margherita.GetName(), margherita.GetCost())

	cheeseTopping := NewCheeseTopping(margherita)
	fmt.Println(cheeseTopping.GetName(), cheeseTopping.GetCost())

	extraVeggies := NewExtraVeggiesPizza(cheeseTopping)
	fmt.Println(extraVeggies.GetName(), extraVeggies.GetCost())

	farmhouse := NewFarmhousePizza()
	fmt.Println(farmhouse.GetName(), farmhouse.GetCost())

	cheeseTopping1 := NewCheeseTopping(farmhouse)
	fmt.Println(cheeseTopping1.GetName(), cheeseTopping1.GetCost())

	extraVeggies1 := NewExtraVeggiesPizza(cheeseTopping1)
	fmt.Println(extraVeggies1.GetName(), extraVeggies1.GetCost())
}
