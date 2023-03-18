package structs

type Circle struct {
	Radius float64
}

func Area(c Circle) float64 {
	pi := 3.141592653589793
	return pi * c.Radius * c.Radius
}

// func Area(rectangle Rectangle) float64 {

// }
