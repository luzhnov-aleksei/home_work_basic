package chessboard

import (
	"strings"
	"testing"
)

func TestChessBoard(t *testing.T) {
	tests := []struct {
		width       int
		height      int
		expected    string
		expectError bool
	}{
		{8, 8, strings.TrimSpace(`
# # # # 
 # # # #
# # # #
 # # # #
# # # #
 # # # #
# # # #
 # # # #
		`), false},
		{0, 8, "", true},
		{8, 6, strings.TrimSpace(`
# # # # 
 # # # #
# # # #
 # # # #
# # # #
 # # # #
		`), false},
		{-1, 8, "", true},
		{8, -1, "", true},
		{0, 0, "", true},
	}

	for _, tt := range tests {
		_, err := ChessBoard(tt.width, tt.height)

		if (err != nil) != tt.expectError {
			t.Errorf("ChessBoard(%d, %d) ошибка = %v, ожидалась ошибка %v", tt.width, tt.height, err, tt.expectError)
			continue
		}

		if err != nil {
			continue
		}

	}
}
