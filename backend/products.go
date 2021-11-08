package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Stock    int    `json:"stock"`
	ImageURL string `json:"image_url"`
}

type ProductsResponse struct {
	Products []*Product `json:"products"`
}

var products = []*Product{
	{
		ID:       "SHOP-0011",
		Name:     "Milk",
		Price:    10,
		Stock:    20,
		ImageURL: "https://images.unsplash.com/photo-1576186726115-4d51596775d1?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=1000&q=80",
	},
	{
		ID:       "SHOP-0012",
		Name:     "Soda",
		Price:    11,
		Stock:    5,
		ImageURL: "https://img-4.linternaute.com/mKXTbLlU0h1MoT6HRZy_FDyK7CI=/900x/smart/a8506dcf71b64b78bb5a66c7ce9cc05c/ccmcms-linternaute/13497750.jpg",
	},
	{
		ID:       "SHOP-0021",
		Name:     "Broccoli",
		Price:    5,
		Stock:    10,
		ImageURL: "https://img-3.journaldesfemmes.fr/mgGEPie3NTYkwkOoGHOLs9AP9ZI=/1240x/smart/c74c02c06f094c709fe3478c583d0a4a/ccmcms-jdf/10659093.jpg",
	},
	{
		ID:       "SHOP-0022",
		Name:     "Carrots",
		Price:    5,
		Stock:    10,
		ImageURL: "https://images.unsplash.com/photo-1447175008436-054170c2e979?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1000&q=80",
	},
	{
		ID:       "SHOP-0031",
		Name:     "Eggs",
		Price:    8,
		Stock:    20,
		ImageURL: "https://images.unsplash.com/photo-1587486913049-53fc88980cfc?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=1000&q=80",
	},
	{
		ID:       "SHOP-0041",
		Name:     "Cheese",
		Price:    20,
		Stock:    2,
		ImageURL: "https://sc04.alicdn.com/kf/U3cddaf0f9eb7489893534e06ab0a856ab.jpeg",
	},
}

// Handler replies with the products response to any request
func Handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	b, err := json.Marshal(ProductsResponse{products})

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func main() {

	fmt.Println("Starting server at port 8080. URL: http://localhost:8080/")

    http.Handle("/", http.HandlerFunc(Handler))
	http.ListenAndServe(":8080", nil)
}
