package dao

import (
	"fmt"
	"week2/model"

	"github.com/jmoiron/sqlx"
)

type MysqlRepo struct {
	dbConn *sqlx.DB
}

func NewMysqlRepo(addr string, user string, passwd string, dbname string) DBRepository {
	db, err := sqlx.Connect("mysql", "")
	if err != nil {
		panic("init mysql failed")
	}

	return &MysqlRepo{
		dbConn: db,
	}
}

func (m *MysqlRepo) QueryUser(userID int) (user *model.UserInfo, err error) {
	user = &model.UserInfo{}
	query := "select * from user where user_id=?"
	err = m.dbConn.Select(user, query, userID)
	if err != nil {
		return nil, fmt.Errorf("query user failed: %w", err)
	}
	return
}
