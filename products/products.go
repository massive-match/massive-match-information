package products

import (
	"fmt"
	"github.com/massive-match/massive-match-information/products/eda"
)

func LoadProducts() {
	path := "data/products/skuPricesWebMkp.csv"

	data := eda.ReadAllProducts(path)
	fmt.Println(data)

	//url := "http://localhost:3000/v1/products/mkp"
	//b, _ := json.Marshal(&data)
	//payload := bytes.NewReader(b)
	//
	//req, _ := http.NewRequest("POST", url, payload)
	//res, _ := http.DefaultClient.Do(req)
	//
	//defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)
	//
	//fmt.Println(res)
	//fmt.Println(string(body))

}
