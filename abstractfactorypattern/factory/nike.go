package factory

import "com.systemdesignpatterns.go.abstractfactorypattern/model"

type Nike struct {
}

func (n *Nike) CreateShirt() model.IShirt {
	return &model.NikeShirt{
		Shirt: model.Shirt{
			Color: "Blue",
			Size:  "M",
			Price: 100,
		},
	}
}

func (n *Nike) CreateShoe() model.IShoe {
	return &model.NikeShoe{
		Shoe: model.Shoe{
			Size:  10,
			Price: 100,
		},
	}
}
