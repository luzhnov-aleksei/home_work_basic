package init

import (
	"fmt"

	"github.com/luzhnov-aleksei/hw02_fix_app/printer"
	"github.com/luzhnov-aleksei/hw02_fix_app/reader"
	"github.com/luzhnov-aleksei/hw02_fix_app/types"
)

func init() {
	path := "data.json"

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path)
	if err != nil {
		fmt.Print(err)
	}

	printer.PrintStaff(staff)
}
