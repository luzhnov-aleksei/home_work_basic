package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		arr         []int
		target      int
		expected    int
		expectError string
	}{
		{[]int{0, 1, 2, 3, 5}, 4, -1, "значение не найдено"},
		{[]int{0, 1, 2, 3, 4}, 0, 0, ""},
		{[]int{0, 1, 2, 3, 4}, 4, 4, ""},
		{[]int{6, 8, 10, 23, 44}, 100, -1, "значение вне диапазона данных"},
	}
	for _, tt := range tests {
		t.Run("BinarySearch", func(t *testing.T) {
			result, err := BinarySearch(tt.arr, tt.target)
			if (err != nil) && err.Error() != tt.expectError {
				t.Errorf("BinarySearch(%v, %d) ошибка = %v, ожидалась ошибка %v", tt.arr, tt.target, err, tt.expectError)
			}
			if result != tt.expected {
				t.Errorf("BinarySearch(%v, %d) = %d, ожидалось %d", tt.arr, tt.target, result, tt.expected)
			}
		})
	}
}
