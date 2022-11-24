package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	Id             int    `json:"id"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerunit"`
	QuantityOnHand int    `json:"quantityOnHand"`
	Name           string `json:"name"`
}

var products []Product

func init() {
	productsJson := `[
		{
			"Id": 1,
			"manufacturer": "Johns-Jenkins",
			"sku": "p5z343vdS",
			"upc": "939581000000",
			"pricePerUnit": "497.45",
			"quantityOnHand": 9703,
			"name": "sticky note"
		},
		{
			"Id": 2,
			"manufacturer": "Hessel, Schimmel and Feeney",
			"sku": "i7v300kmx",
			"upc": "740979000000",
			"pricePerUnit": "282.29",
			"quantityOnHand": 9217,
			"name": "leg warmers"
		},
		{
			"Id": 3,
			"manufacturer": "Swaniawski, Bartoletti and Bruen",
			"sku": "q0L657ys7",
			"upc": "111730000000",
			"pricePerUnit": "436.26",
			"quantityOnHand": 5905,
			"name": "lamp shade"
		}
	]`

	err := json.Unmarshal([]byte(productsJson), &products)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/products", productHandler)
	http.HandleFunc("/products/", productsHandler)
	http.ListenAndServe(":3000", nil)
}

func getNextId() int {
	highestId := -1
	for _, product := range products {
		if highestId < product.Id {
			highestId = product.Id
		}
	}
	return highestId + 1
}

func findProductById(productId int) (*Product, int) {
	for i, product := range products {
		if product.Id == productId {
			return &product, i
		}
	}
	return nil, 0
}

func producstHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, "products/")
	prductId, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	product, index := findProductById(prductId)
	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		encodeResponseAsJSON(products, w)
		w.Header().Set("Content-Type", "application/json")

	case http.MethodPost:
		var newProduct Product
		newProduct, err := parseRequest(r)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if newProduct.Id != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newProduct.Id = getNextId()
		products = append(products, newProduct)
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func parseRequest(r *http.Request) (Product, error) {
	dec := json.NewDecoder(r.Body)
	var newProduct Product
	err := dec.Decode(&newProduct)
	if err != nil {
		return Product{}, err
	}
	return newProduct, nil
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
