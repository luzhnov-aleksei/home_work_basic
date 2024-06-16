package chessboard

import (
	"errors"
	"strings"
)

func ChessBoard(width, height int) (string, error) {
	if width <= 0 || height <= 0 {
		return "", errors.New("width and height must be greater than zero")
	}

	var board strings.Builder
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if (i+j)%2 == 0 {
				board.WriteRune('#') // Black square representation
			} else {
				board.WriteRune(' ') // White square representation
			}
		}
		board.WriteString("\n")
	}
	return board.String(), nil
}
