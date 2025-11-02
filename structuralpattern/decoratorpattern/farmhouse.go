package main

type FarmhousePizza struct {
	Pizza
}

func NewFarmhousePizza() IPizza {
	return &FarmhousePizza{
		Pizza: Pizza{
			Description: "Farmhouse",
			Cost:        12.0,
		},
	}
}
