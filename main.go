package main

import (
	"io/ioutil"
	"encoding/json"
	"log"
	"net/http"
)

type Product struct {
	ProductID      int
	Manufacturer   string
	Sku            string
	Upc            string
	PricePerUnit   string
	QuantityOnHand int
	ProductName    string
}

var productList []Product

func init() {
	productsJSON := `[
		{
			"ProductID": 1,
			"Manufacturer": "Johns-Jenkins",
			"Sku": "p5z",
			"Upc": "1234",
			"PricePerUnit": "459.34",
			"QuantityOnHand": 9703,
			"ProductName": "foo"
		},
		{
			"ProductID": 2,
			"Manufacturer": "H, S, and F",
			"Sku": "i7v",
			"Upc": "74097",
			"PricePerUnit": "282.29",
			"QuantityOnHand": 9217,
			"ProductName": "leg warmers"
		},
		{
			"ProductID": 3,
			"Manufacturer": "S, B and B",
			"Sku": "q0L",
			"Upc": "11173",
			"PricePerUnit": "436.34",
			"QuantityOnHand": 5905,
			"ProductName": "lamp shade"
		}
	]`
	err := json.Unmarshal([]byte(productsJSON), &productList)
	if err != nil {
		log.Fatal(err)
	}
}

func getNextID() int {
	highestID := -1
	for _, product := range productList {
		if highestID < product.ProductID {
			highestID = product.ProductID
		}
	}
	return highestID + 1
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productsJson, err := json.Marshal(productList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productsJson)
	case http.MethodPost:
		var newProduct Product
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &newProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if newProduct.ProductID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newProduct.ProductID = getNextID()
		productList = append(productList, newProduct)
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func main() {
	http.HandleFunc("/products", productsHandler)
	http.ListenAndServe(":5000", nil)
}
