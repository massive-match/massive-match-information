package eda

import (
	"github.com/tealeg/xlsx"
)

type Matrix struct {
	JotaID     string `json:"jota_id" xlsx:"0"`
	Attr       string `json:"attr" xlsx:"1"`
	Type       string `json:"type" xlsx:"2"`
	OptionType string `json:"option_type" xlsx:"3"`
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
