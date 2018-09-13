package eda

import "github.com/tealeg/xlsx"

type MatchCategories struct {
	MclL1ID  string `json:"mcl_l1_id" xlsx:"0"`
	MclL1    string `json:"mcl_l1" xlsx:"1"`
	MclL2ID  string `json:"mcl_l2_id" xlsx:"2"`
	MclL2    string `json:"mcl_l2" xlsx:"3"`
	MclL3ID  string `json:"mcl_l3_id" xlsx:"4"`
	MclL3    string `json:"mcl_l3" xlsx:"5"`
	MclL4ID  string `json:"mcl_l4_id" xlsx:"6"`
	MclL4    string `json:"mcl_l4" xlsx:"7"`
	MclL5ID  string `json:"mcl_l5_id" xlsx:"8"`
	MclL5    string `json:"mcl_l5" xlsx:"9"`
	FcomL1ID string `json:"fcom_l1_id" xlsx:"10"`
	FcomL1   string `json:"fcom_l1" xlsx:"11"`
	FcomL2ID string `json:"fcom_l2_id" xlsx:"12"`
	FcomL2   string `json:"fcom_l2" xlsx:"13"`
	FcomL3ID string `json:"fcom_l3_id" xlsx:"14"`
	FcomL3   string `json:"fcom_l3" xlsx:"15"`
	FcomL4ID string `json:"fcom_l4_id" xlsx:"16"`
	FcomL4   string `json:"fcom_l4" xlsx:"17"`
}

func ReadMatchCategories(path string) []MatchCategories {
	var matchCategories []MatchCategories

	if xlFile, err := xlsx.OpenFile(path); err != nil {
		return matchCategories
	} else {
		for i, sheet := range xlFile.Sheets {
			if i == 0 {
				matchCategories = make([]MatchCategories, len(sheet.Rows))
				for j, row := range sheet.Rows {
					row.ReadStruct(&matchCategories[j])
				}
			}
		}
		return matchCategories
	}
}

