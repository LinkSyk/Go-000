package usecase

import "week2/model"

type UseCaseInterface interface {
	QueryUser(userID int) (*model.UserInfo, error)
}
