package chessboard

import (
	"errors"
	"fmt"
)

func ChessBoard(width, height int) error {
	if width <= 0 {
		return errors.New("некорректное значение ширины поля, введите положительное целое число")
	}
	if height <= 0 {
		return errors.New("некорректное значение высоты поля, введите положительное целое число")
	}

	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			if j%2 == 0 {
				fmt.Print(" #")
			} else {
				fmt.Print("# ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	return nil
}
