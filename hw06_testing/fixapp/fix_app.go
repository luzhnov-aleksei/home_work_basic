package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/luzhnov-aleksei/hw02_fix_app/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	defer f.Close()

	byteChar, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	if len(byteChar) == 0 {
		return []types.Employee{}, nil
	}

	var data []types.Employee

	err = json.Unmarshal(byteChar, &data)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	return data, nil
}
