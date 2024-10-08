package shapes

import (
	"errors"
	"math"
)

type shape interface {
	area() (float64, error)
}

type circle struct {
	radius float64
}

type rectangle struct {
	width  float64
	height float64
}

type triangle struct {
	base   float64
	height float64
}

func (c circle) area() (float64, error) {
	if c.radius < 0 {
		return 0, errors.New("радиус не может быть отрицательным")
	}
	return math.Pi * math.Pow(c.radius, 2), nil
}

func (r rectangle) area() (float64, error) {
	if r.width < 0 || r.height < 0 {
		return 0, errors.New("ширина и высота не могут быть отрицательными")
	}
	return r.width * r.height, nil
}

func (t triangle) area() (float64, error) {
	if t.base < 0 || t.height < 0 {
		return 0, errors.New("основание и высота не могут быть отрицательными")
	}
	return t.base * t.height / 2, nil
}

func calculateArea(data any) (float64, error) {
	s, ok := data.(shape)
	if !ok {
		return 0.0, errors.New("неизвестный тип фигуры")
	}
	return s.area()
}
