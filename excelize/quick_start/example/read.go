package example

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func Read(fileName string) error {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		return err
	}

	defer func() {
		if err = f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// 获取工作表中指定单元格的值
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		return err
	}

	fmt.Println(cell)

	// 获取Sheet1 上的所有单元格
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return err
	}

	for _, row := range rows {
		for _, collCell := range row {
			fmt.Print(collCell, "\t")
		}
		fmt.Println()
	}

	return nil
}
