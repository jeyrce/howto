package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// https://xuri.me/excelize/zh-hans/
func init() {

}

func main() {
	file := excelize.NewFile()
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	sheet, err := file.NewSheet("sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	file.SetCellValue("sheet1", "A1", "jeyrce")
	file.SetActiveSheet(sheet)
	if err := file.SaveAs("./test.xlsx"); err != nil {
		panic(err)
	}
}
