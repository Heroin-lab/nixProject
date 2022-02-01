package models

type Suppliers struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	Type         string `json:"type"`
	Working_time string `json:"working_time"`
}

type SuppliersForSelect struct {
	Title        string
	Type         string
	Working_time string
}
