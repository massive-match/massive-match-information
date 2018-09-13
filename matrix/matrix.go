package matrix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/massive-match/massive-match-information/matrix/eda"
)

func LoadMatrix() {
	path := "data/matrix/matrix-svl.xlsx"

	var matrices []eda.Matrix
	matrices = eda.ReadMatrix(path)

	url := "http://13.66.196.10/api/v1/internal/categories/matrix"
	b, _ := json.Marshal(&matrices)
	payload := bytes.NewReader(b)

	req, _ := http.NewRequest("POST", url, payload)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
