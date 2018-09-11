package eda

import (
	"github.com/tealeg/xlsx"
)

type Matrix struct {
	Id         int    `xlsx:"0"`
	Attr       string `xlsx:"1"`
	Obs        string `xlsx:"2"`
	Obligatory string `xlsx:"3"`
}

func ReadMatrix(path string) []Matrix {
	var matrices []Matrix

	if xlFile, err := xlsx.OpenFile(path); err != nil {
		return matrices
	} else {
		for i, sheet := range xlFile.Sheets {
			if i == 0 {
				matrices = make([]Matrix, len(sheet.Rows))
				for j, row := range sheet.Rows {
					row.ReadStruct(&matrices[j])
				}
			}
		}
		return matrices
	}
}
