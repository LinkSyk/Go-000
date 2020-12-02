package handler

import (
	"log"
	"net/http"
	"week2/model"
	"week2/usecase"

	echo "github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	usecase usecase.UseCaseInterface
}

func NewHandler(uc usecase.UseCaseInterface) *Handler {
	return &Handler{
		usecase: uc,
	}
}

func (h *Handler) Run() {
	e := echo.New()
	// add auth middleware
	// .....
	e.POST("/create", h.queryUser)
	log.Fatal(e.Start(":8080"))
}

func (h *Handler) queryUser(c echo.Context) error {

	ret := struct {
		ErrCode int             `json:"err_code"`
		user    *model.UserInfo `json:"user"`
	}{
		1001, // 1000 表示正常, 1001 表示用户不存在, 1002 表示其他错误
		nil,
	}

	var info struct {
		UserID int `json:"user_id"`
	}

	if err := c.Bind(&info); err != nil {
		logrus.Infof("invalid param: %v", err)
		ret.ErrCode = 1002
		return c.JSON(http.StatusOK, &ret)
	}

	user, err := h.usecase.QueryUser(info.UserID)
	if err != nil {
		// if another error happen, log the error and return err
		logrus.Infof("query user failed: %s", err)
		return err
	}

	return c.JSON(http.StatusOK, user)
}
