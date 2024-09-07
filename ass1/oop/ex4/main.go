package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name     string `json:"name"`
	Price    uint   `json:"price"`
	Quantity uint   `json:"quantity"`
}

func (product Product) toJson() string {
	data, err := json.Marshal(product)

	if err != nil {
		fmt.Println("Error in converting to json:", err)
		return ""
	}
	return string(data)
}

func decodeJson(jsonData string) {
	var product Product
	err := json.Unmarshal([]byte(jsonData), &product)

	if err != nil {
		fmt.Println("Error in decoding a json:", err)
		return
	}
	fmt.Println("Product:", product)
}

func main() {
	product := Product{Name: "Laptop", Price: 1500, Quantity: 100}
	fmt.Println("Product:", product)

	var jsonData = product.toJson()
	fmt.Println("json format:", jsonData)

	decodeJson(jsonData)
}
