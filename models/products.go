package models

type Products struct {
	Id           int     `json:"id"`
	Product_name string  `json:"product_name"`
	Type_id      int     `json:"type_id"`
	Price        float64 `json:"price"`
	Img          string  `json:"img"`
	Supplier_id  int     `json:"supplier_id"`
}

type ForSelectProducts struct {
	Id             string
	Product_name   string
	Prod_type_name string
	Price          float64
	Img            string
	Supplier       string
	Quantity       int
}

//
//type Product struct {
//	Id          int      `json:"id"`
//	Name        string   `json:"name"`
//	Price       float64  `json:"price"`
//	Image       string   `json:"image"`
//	Type        string   `json:"type"`
//	Ingredients []string `json:"ingredients"`
//}
