package model

type UserInfo struct {
	UserID int    `json:"user_id" db:"user_id"`
	Name   string `json:"name" db:"name"`
	Age    int    `json:"age" db:"age"`
}
