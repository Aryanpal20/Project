package datamodel

type Data struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Rating      string `json:"rating"`
	Price       string `json:"price"`
	Category_ID int    `json:"category_id"`
}
