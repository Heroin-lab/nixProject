package models

type Products struct {
	Id           int    `json:"id"`
	Product_name string `json:"product_name"`
	Category_id  int    `json:"category_id"`
	Price        string `json:"price"`
	Prod_desc    string `json:"prod_desc"`
	Amount_left  int    `json:"amount_left"`
	Supplier_id  int    `json:"supplier_id"`
}

type ForSelectProducts struct {
	Id            string
	Product_name  string
	Category_name string
	Price         string
	Prod_desc     string
	Amount_left   int
	Title         string
	Quantity      int
}
