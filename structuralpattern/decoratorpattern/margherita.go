package main

type MargheritaPizza struct {
	Pizza
}

func NewMargheritaPizza() IPizza {
	return &MargheritaPizza{
		Pizza: Pizza{
			Description: "Margherita",
			Cost:        10.0,
		},
	}
}
