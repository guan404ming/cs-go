package models

type Category struct {
	Name     string   `json:"name"`
	Listings []string `json:"listings"` // 該類別下的商品ID列表
}
