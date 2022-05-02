package models

type Suppliers struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Image string `json:"image"`
	WorkingHours
	Menu []Products `json:"menu"`
}

type WorkingHours struct {
	Opening string `json:"opening"`
	Closing string `json:"closing"`
}
