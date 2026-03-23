package products

import "time"

type ProductsModel struct {
	ID	int8 `json:"id"`
	Name string `json:"name"`
	Createat time.Time `json:"createat"`
}