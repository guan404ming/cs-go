package models

type User struct {
	Username string   `json:"username"`
	Listings []string `json:"listings"` // 存儲該用戶的商品ID列表
}
