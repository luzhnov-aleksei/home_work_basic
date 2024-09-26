package shapes

import (
	"testing"
)

func TestCalculateArea(t *testing.T) {
	tests := []struct {
		name       string
		shapeData  interface{}
		expected   float64
		errMessage string
	}{
		{"Круг", circle{radius: 3}, 28.274333882308138, ""},
		{"Круг_НеверныйРадиус", circle{radius: -3}, 0, "радиус не может быть отрицательным"},
		{"Прямоугольник", rectangle{width: 4, height: 5}, 20, ""},
		{
			"Прямоугольник_НевернаяШирина",
			rectangle{width: -4, height: 5},
			0,
			"ширина и высота не могут быть отрицательными",
		},
		{"Треугольник", triangle{base: 4, height: 3}, 6, ""},
		{
			"Треугольник_НеверноеОснование",
			triangle{base: -4, height: 3},
			0,
			"основание и высота не могут быть отрицательными",
		},
		{"НеизвестнаяФигура", struct{}{}, 0, "неизвестный тип фигуры"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			area, err := calculateArea(tt.shapeData)
			if err != nil {
				if err.Error() != tt.errMessage {
					t.Errorf("получена ошибка %v, ожидалась ошибка %v", err, tt.errMessage)
				}
			} else {
				if area != tt.expected {
					t.Errorf("получена площадь %f, ожидалась площадь %f", area, tt.expected)
				}
			}
		})
	}
}
