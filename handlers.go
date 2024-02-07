package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

var productsBody = new(ProductsResponse)
var productsSlice []Product
var tmpl *template.Template

func HandleGet(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, productsBody.Products)
}

func HandleProductUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	quantity := r.PostFormValue("Quantity")

	if len(quantity) == 0 || len(productID) == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	Quantity, err := strconv.Atoi(quantity)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	ProductID, err := strconv.Atoi(productID)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var target Product

	for i := range productsSlice {
		if productsSlice[i].Id == ProductID {
			productsSlice[i].Price = (productsSlice[i].Price / productsSlice[i].Quantity) * Quantity
			productsSlice[i].Quantity = Quantity
			target = productsSlice[i]
			break
		}
	}

	if target == (Product{}) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	err = tmpl.ExecuteTemplate(w, "cart-table-row", target)
	if err != nil {
		fmt.Println(err)
	}
}

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	var results []Product
	searchQuery := strings.TrimSpace(r.URL.Query().Get("q"))

	if len(searchQuery) == 0 {
		w.Write(nil)
		return
	}

	productsSliceMap := make(map[int]Product)

	for _, product := range productsSlice {
		productsSliceMap[product.Id] = product
	}

	for i := range productsBody.Products {
		name := productsBody.Products[i].Title

		if strings.Contains(strings.ToLower(name), strings.ToLower(searchQuery)) {
			foundProduct := productsBody.Products[i]
			if _, exists := productsSliceMap[foundProduct.Id]; exists {
				foundProduct.Disabled = true
			}
			results = append(results, foundProduct)
		}
	}

	err := tmpl.ExecuteTemplate(w, "search-results", results)
	if err != nil {
		fmt.Println(err)
	}
}

func HandleProductDelete(w http.ResponseWriter, r *http.Request) {
	var products []Product

	vars := mux.Vars(r)
	productID := vars["id"]

	if len(productID) == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	productId, err := strconv.Atoi(productID)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	for i := range productsSlice {
		if productsSlice[i].Id != productId {
			products = append(products, productsSlice[i])
		}
	}

	productsSlice = products
	w.Write(nil)
}

func HandleProductInsert(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	var foundProduct Product

	if len(productID) == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	productId, err := strconv.Atoi(productID)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	productsSliceMap := make(map[int]Product)

	for _, product := range productsSlice {
		productsSliceMap[product.Id] = product
	}

	for i := range productsBody.Products {
		if productsBody.Products[i].Id == productId {
			foundProduct = productsBody.Products[i]
			break
		}
	}

	if foundProduct == (Product{}) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if _, exists := productsSliceMap[foundProduct.Id]; exists {
		http.Error(w, "Product already in the cart", http.StatusBadRequest)
		return
	}

	productsSlice = append(productsSlice, foundProduct)

	err = tmpl.ExecuteTemplate(w, "product-info-popup", foundProduct)
	if err != nil {
		fmt.Println(err)
	}
}
