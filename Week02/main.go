package main

import (
	"week2/dao"
	"week2/handler"
	"week2/usecase"
)

func main() {
	mysqlRepo := dao.NewMysqlRepo("127.0.0.1", "readyonly", "xxxx", "user")
	usecase := usecase.NewUserManager(mysqlRepo)
	h := handler.NewHandler(usecase)
	h.Run()
}
