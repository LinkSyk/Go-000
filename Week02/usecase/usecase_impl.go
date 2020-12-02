package usecase

import (
	"database/sql"
	"errors"
	"week2/dao"
	"week2/model"
)

type UserManger struct {
	dbrepo dao.DBRepository
}

func NewUserManager(db dao.DBRepository) UseCaseInterface {
	return &UserManger{
		dbrepo: db,
	}
}

func (um *UserManger) QueryUser(userID int) (user *model.UserInfo, err error) {
	user, err = um.dbrepo.QueryUser(userID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return
}
