package factory

import "com.systemdesignpatterns.go.abstractfactorypattern/model"

type AbstractFactory interface {
	CreateShirt() model.IShirt
	CreateShoe() model.IShoe
}

func GetFactory(brand string) AbstractFactory {
	switch brand {
	case "Adidas":
		return &Adidas{}
	case "Nike":
		return &Nike{}
	default:
		return nil
	}
}
