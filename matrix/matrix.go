package matrix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dwladdimiroc/massive-match-information/matrix/eda"
)

func LoadProducts() {
	path := "data/MatricesCorp.xlsx"

	var matrices []eda.Matrix
	matrices = eda.ReadMatrix(path)
	//fmt.Println(matrices)

	var list []eda.List
	list = eda.ReadList(path)
	//fmt.Println(list)

	var relations []eda.Relation
	relations = eda.ReadJotas(path)
	//fmt.Println(relations)

	data := eda.MatchCategories(matrices, list, relations)

	url := "http://localhost:3000/v1/categories/matrix"
	b, _ := json.Marshal(&data)
	payload := bytes.NewReader(b)

	req, _ := http.NewRequest("POST", url, payload)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
