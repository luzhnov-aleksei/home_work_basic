package main

import (
	"errors"
)

func BinarySearch(arr []int, target int) (int, error) {
	left := 0
	right := len(arr) - 1
	if target < arr[left] || target > arr[right] {
		return -1, errors.New("значение вне диапазона данных")
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
	return -1, errors.New("значение не найдено")
}
