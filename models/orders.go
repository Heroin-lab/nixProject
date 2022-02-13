package models

type Order struct {
	Id          string  `json:"id"`
	Paid_status bool    `json:"paid_Status"`
	Address     string  `json:"address"`
	Price       float64 `json:"price"`
	User_id     string  `json:"user_Id"`
	ProductArr  []ForSelectProducts
}

type OrderUID struct {
	User_id string `json:"user_Id"`
}

type OrderForInsert struct {
	Id          string  `json:"id"`
	Paid_status bool    `json:"paid_Status"`
	Address     string  `json:"address"`
	Price       float64 `json:"price"`
	User_id     string  `json:"user_Id"`
	ProductArr  string  `json:"productArr"`
}
