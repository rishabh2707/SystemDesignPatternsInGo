package main

type IPizza interface {
	GetName() string
	GetCost() float64
}

type Pizza struct {
	Description string
	Cost        float64
}

func (p *Pizza) GetName() string {
	return p.Description
}

func (p *Pizza) GetCost() float64 {
	return p.Cost
}

func (p *Pizza) SetName(name string) {
	p.Description = name
}

func (p *Pizza) SetCost(cost float64) {
	p.Cost = cost
}
