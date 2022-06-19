package example

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func Create(name string) error {
	f := excelize.NewFile()
	index := f.NewSheet("Sheet2")
	err := f.SetCellValue("Sheet2", "A2", "Hello excelize")
	if err != nil {
		return err
	}
	err = f.SetCellValue("Sheet1", "B2", 100)
	if err != nil {
		return err
	}

	f.SetActiveSheet(index)

	if err = f.SaveAs(name); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
