package repo

import (
	"fmt"
	"time"
	"week04/internal/model"

	"github.com/jmoiron/sqlx"
)

type DBRepository interface {
	CreateUser(user *model.User) error
	DelUser(user *model.SimpleUser) error
}

type MysqlDBImpl struct {
	db *sqlx.DB
}

func ConnectMysql(user, pass, host, dbname string, port int16) (*sqlx.DB, error) {
	conn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, dbname)
	return sqlx.Connect("mysql", conn)
}

func NewMysqlDBImpl(db *sqlx.DB) DBRepository {
	return &MysqlDBImpl{
		db: db,
	}
}

func (m *MysqlDBImpl) CreateUser(user *model.User) error {
	query := "insert into User(nick_name, gender, agem phone, delete_at) values(?,?,?,?,NULL)"
	_, err := m.db.Exec(query, user.NickName, user.Gender, user.Age, user.Phone)
	if err != nil {
		// 返回业务错误 ErrCreateUser
		return fmt.Errorf("err: %w, sql: %s, args: %v", model.ErrCreateUser, query, user)
	}
	return nil
}

func (m *MysqlDBImpl) DelUser(user *model.SimpleUser) error {
	query := "update User set delete_at = ? where id=?"
	_, err := m.db.Exec(query, time.Now(), user.UserID)
	if err != nil {
		// 返回业务错误 ErrDelUser
		return fmt.Errorf("err: %w, sql: %s, args: %v", model.ErrDelUser, query, user)
	}
	return nil
}
