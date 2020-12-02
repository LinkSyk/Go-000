package dao

import "week2/model"

type DBRepository interface {
	QueryUser(userID int) (*model.UserInfo, error)
}
