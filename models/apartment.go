package models

type Apartment struct {
	Id int `json:"id"`
    Title string `json:"title"`
	Description string `json:"description"`
	Price float32 `json:"price"`
	Rental bool `json:"Rental"`
	Available bool `json:"Open"`
}
