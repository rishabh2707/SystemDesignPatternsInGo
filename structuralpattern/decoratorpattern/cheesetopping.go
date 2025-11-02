package main

type CheeseTopping struct {
	IPizza
}

func (t *CheeseTopping) GetName() string {
	return t.IPizza.GetName() + " with " + " extra cheese"
}

func (t *CheeseTopping) GetCost() float64 {
	return t.IPizza.GetCost() + 2
}

func NewCheeseTopping(pizza IPizza) IPizza {
	return &CheeseTopping{
		IPizza: pizza,
	}
}
