package eda

import "github.com/tealeg/xlsx"

type List struct {
	Attr  string `xlsx:"0"`
	Value string `xlsx:"1"`
}

func ReadList(path string) []List {
	var list []List

	if xlFile, err := xlsx.OpenFile(path); err != nil {
		return list
	} else {
		for i, sheet := range xlFile.Sheets {
			if i == 1 {
				list = make([]List, len(sheet.Rows))
				for j, row := range sheet.Rows {
					row.ReadStruct(&list[j])
				}
			}
		}
		return list
	}
}
