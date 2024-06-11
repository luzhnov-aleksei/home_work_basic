package chessboard

import (
	"testing"
)

func TestChessBoard(t *testing.T) {
	tests := []struct {
		width       int
		height      int
		expectError bool
	}{
		{8, 8, false},
		{0, 8, true},
		{8, 6, false},
		{-1, 8, true},
		{8, -1, true},
		{0, 0, true},
	}

	for _, tt := range tests {
		err := ChessBoard(tt.width, tt.height)
		if (err != nil) != tt.expectError {
			t.Errorf("ChessBoard(%d, %d) error = %v, expectError %v", tt.width, tt.height, err, tt.expectError)
		}
	}
}
