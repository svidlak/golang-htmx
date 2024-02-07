package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func populateProducts() {
	resp, err := http.Get("https://dummyjson.com/products")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(productsBody)
	for i := range productsBody.Products {
		productsBody.Products[i].Quantity = 1
	}

	productsSlice = productsBody.Products[:4]
}

func main() {
	r := mux.NewRouter()

	populateProducts()
	initAndBumdleTemplates()

	http.Handle("/", r)
	r.HandleFunc("/", HandleGet).Methods("GET")
	r.HandleFunc("/search", HandleSearch).Methods("GET")
	r.HandleFunc("/product/{id}", HandleProductUpdate).Methods("PUT")
	r.HandleFunc("/product/{id}", HandleProductDelete).Methods("DELETE")
	r.HandleFunc("/product/{id}", HandleProductInsert).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", nil))
}
