package main

type ExtraVeggiesPizza struct {
	IPizza
}

func (t *ExtraVeggiesPizza) GetName() string {
	return t.IPizza.GetName() + " with " + " extra veggies"
}

func (t *ExtraVeggiesPizza) GetCost() float64 {
	return t.IPizza.GetCost() + 3
}

func NewExtraVeggiesPizza(pizza IPizza) IPizza {
	return &ExtraVeggiesPizza{
		IPizza: pizza,
	}
}
