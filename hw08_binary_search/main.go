package main

import (
	"errors"
	"fmt"
)

func binarySearch(arr []int, target int) (int, error) {
	left := 0
	right := len(arr) - 1
	if target < arr[left] || target > arr[right] {
		return 0, errors.New("значение вне диапазона данных")
	}
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid, nil
		}
		if target < arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return 0, errors.New("значение не найдено")
}

func main() {
	arr := []int{2, 3, 4, 5, 6, 7, 9, 10, 12}
	a, err := binarySearch(arr, 11)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Найдено значение %d по индексу %d \n", arr[a], a)
	}
}
