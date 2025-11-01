package factory

import "com.systemdesignpatterns.go.factorypattern/model"

func GetShape(shapeType string) model.IShape {
	switch shapeType {
	case "rectangle":
		return model.Rectangle{Width: 10, Height: 20}
	case "circle":
		return model.Circle{Radius: 10}
	case "square":
		return model.Square{Side: 10}
	}
	return nil
}
