package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
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

func calculateArea(s shape) (float64, error) {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if !v.Type().Implements(reflect.TypeOf((*shape)(nil)).Elem()) {
		return 0, errors.New("некорректный тип данных, ожидается тип, реализующий интерфейс Shape")
	}
	return s.area()
}

func main() {
	c1 := circle{radius: 5}
	r1 := rectangle{width: 10, height: 5}
	t1 := triangle{base: 8, height: 6}

	if area, err := calculateArea(c1); err != nil {
		fmt.Printf("Ошибка при вычислении площади круга: %v\n", err)
	} else {
		fmt.Printf("Круг: радиус %g, площадь: %g\n", c1.radius, area)
	}
	if area, err := calculateArea(r1); err != nil {
		fmt.Printf("Ошибка при вычислении площади прямоугольника: %v\n", err)
	} else {
		fmt.Printf("Прямоугольник: ширина %g, высота %g, площадь: %g\n", r1.width, r1.height, area)
	}

	if area, err := calculateArea(t1); err != nil {
		fmt.Printf("Ошибка при вычислении площади треугольника: %v\n", err)
	} else {
		fmt.Printf("Треугольник: основание %g, высота %g, площадь: %g\n", t1.base, t1.height, area)
	}
}
