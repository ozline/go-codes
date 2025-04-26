package code

// Excelize 读写 Excel 的 demo

import (
	"fmt"
	"time"

	"github.com/xuri/excelize/v2"
)

func RunExcelizeRead() {
	f, err := excelize.OpenFile("example.xlsx")

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	rows, err := f.GetRows("Sheet1")

	if err != nil {
		panic(err)
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}

		for index, colCell := range row {
			fmt.Printf("%v %v\n", index, colCell)
			if index > 7 {
				dateFormat := "2006/01/02 15:04:05"
				dateTime, err := time.Parse(dateFormat, colCell)

				if err != nil {
					panic(err)
				}

				fmt.Println("解析后时间：", dateTime)
			}
		}
		fmt.Println()
	}
}
