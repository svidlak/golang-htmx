package main

type ProductsResponse struct {
	Total    int       `json:"total"`
	Skip     int       `json:"skip"`
	Limit    int       `json:"limit"`
	Products []Product `json:"products"`
}

type Product struct {
	Id                 int     `json:"id"`
	Title              string  `json:"title"`
	Description        string  `json:"description"`
	Price              int     `json:"price"`
	Rating             float32 `json:"rating"`
	Stock              int     `json:"stock"`
	Brand              string  `json:"brand"`
	Category           string  `json:"category"`
	Thumbnail          string  `json:"thumbnail"`
	DiscountPercentage float32 `json:"discountPercentage"`
	Quantity           int
	Disabled           bool
}
