package models

type Products struct {
	Id           int
	Product_name string
	Category_id  int
	Price        string
	Description  string
	Amount_left  int
	Supplier_id  int
}

type SelectProducts struct {
	Id            int    `json:"id"`
	Product_name  string `json:"product_name"`
	Category_name string `json:"category_name"`
	Price         string `json:"price"`
	Description   string `json:"description"`
	Amount_left   int    `json:"amount_left"`
	Title         string `json:"title"`
}
