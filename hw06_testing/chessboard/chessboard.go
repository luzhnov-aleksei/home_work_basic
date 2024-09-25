package chessboard

import (
	"errors"
	"strings"
)

func ChessBoard(width, height int) (string, error) {
	if width <= 0 || height <= 0 {
		return "", errors.New("ширина и высота должны быть больше нуля")
	}

	var board strings.Builder
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if (i+j)%2 == 0 {
				board.WriteRune('#')
			} else {
				board.WriteRune(' ')
			}
		}
		board.WriteString("\n")
	}
	return board.String(), nil
}
