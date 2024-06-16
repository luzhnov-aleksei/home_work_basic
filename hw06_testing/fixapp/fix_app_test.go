package reader

import (
	"os"
	"testing"

	"github.com/luzhnov-aleksei/hw02_fix_app/types"
	"github.com/stretchr/testify/assert"
)

func createTempFile(t *testing.T, content string) *os.File {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	_, err = tmpFile.Write([]byte(content))
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	err = tmpFile.Close()
	if err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	return tmpFile
}

func TestReadJSON(t *testing.T) {
	tests := []struct {
		name          string
		fileContent   string
		expectedData  []types.Employee
		expectedError bool
	}{
		{
			name: "Valid JSON",
			fileContent: `[
				{"userId": 1, "age": 30, "name": "John Doe", "departmentId": 101},
				{"userId": 2, "age": 25, "name": "Jane Smith", "departmentId": 102}
			]`,
			expectedData: []types.Employee{
				{UserID: 1, Age: 30, Name: "John Doe", DepartmentID: 101},
				{UserID: 2, Age: 25, Name: "Jane Smith", DepartmentID: 102},
			},
			expectedError: false,
		},
		{
			name:          "Empty File",
			fileContent:   ``,
			expectedData:  []types.Employee{},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file with the specified content
			tmpFile := createTempFile(t, tt.fileContent)
			defer os.Remove(tmpFile.Name())

			// Call the function under test
			data, err := ReadJSON(tmpFile.Name())

			// Verify the results
			assert.Equal(t, tt.expectedData, data)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
