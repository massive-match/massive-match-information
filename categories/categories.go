package categories

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/massive-match/massive-match-information/categories/eda"
	"io/ioutil"
	"net/http"
)

func LoadRelationCategories() {
	path := "data/categories/mlc-categories.xlsx"

	var matchCategories []eda.MatchCategories
	matchCategories = eda.ReadMatchCategories(path)

	url := "http://13.66.196.10/api/v1/internal/categories/relation"
	b, _ := json.Marshal(&matchCategories)
	payload := bytes.NewReader(b)

	req, _ := http.NewRequest("POST", url, payload)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
