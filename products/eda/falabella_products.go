package eda

import (
	"fmt"
	"github.com/artonge/go-csv-tag"
)

type FalabellaProducts struct {
	Sku              string `csv:"sku"`
	Name             string `csv:"nombre"`
	NameFcom         string `csv:"nombre_fcom"`
	Brand            string `csv:"marca"`
	CatFcom          string `csv:"cat_fcom"`
	NameCatFcom      string `csv:"nombre_cat_fcom"`
	StatePublication string `csv:"estado_publicacion"`
	Price1           string `csv:"precio1"`
	Price2           string `csv:"precio2"`
	Price3           string `csv:"precio3"`
	TypePrice1       string `csv:"tipo_precio1"`
	TypePrice2       string `csv:"tipo_precio2"`
	FcomUrl          string `csv:"fcom_url"`
	Date             string `csv:"date"`
	Origin           string `csv:"origen"`
	JL1              string `csv:"j1id"`
	JL1Name          string `csv:"j1"`
	JL2              string `csv:"j2id"`
	JL2Name          string `csv:"j2"`
	JL3              string `csv:"j3id"`
	JL3Name          string `csv:"j3"`
	JL4              string `csv:"j4id"`
	JL4Name          string `csv:"j4"`
	MKPName          string `csv:"mkp_nombre"`
	Seller           string `csv:"seller"`
}

func ReadAllProducts(path string) []FalabellaProducts {
	var falabellaProducts []FalabellaProducts
	var cfgCsv = csvtag.Config{
		Path:      path,
		Dest:      &falabellaProducts,
		Separator: ';',
	}

	if err := csvtag.Load(cfgCsv); err != nil {
		fmt.Println(err.Error())
		return nil
	} else {
		return falabellaProducts
	}

	return falabellaProducts
}
