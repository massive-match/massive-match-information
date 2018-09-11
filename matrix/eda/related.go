package eda

import "github.com/tealeg/xlsx"

type Relation struct {
	Jota   string `xlsx:"0"`
	Matrix int    `xlsx:"1"`
}

func ReadJotas(path string) []Relation {
	var relation []Relation

	if xlFile, err := xlsx.OpenFile(path); err != nil {
		return relation
	} else {
		for i, sheet := range xlFile.Sheets {
			if i == 2 {
				relation = make([]Relation, len(sheet.Rows))
				for j, row := range sheet.Rows {
					row.ReadStruct(&relation[j])
				}
			}
		}
		return relation
	}
}
