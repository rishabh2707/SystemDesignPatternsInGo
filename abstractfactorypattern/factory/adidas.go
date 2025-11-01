package factory

import "com.systemdesignpatterns.go.abstractfactorypattern/model"

type Adidas struct {
}

func (a *Adidas) CreateShirt() model.IShirt {
	return &model.AdidasShirt{
		Shirt: model.Shirt{
			Color: "Red",
			Size:  "L",
			Price: 100,
		},
	}
}

func (a *Adidas) CreateShoe() model.IShoe {
	return &model.AdidasShoe{
		Shoe: model.Shoe{
			Size:  10,
			Price: 100,
		},
	}
}
