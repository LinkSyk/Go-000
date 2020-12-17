package service

import (
	"context"

	account "week04/api/account/v1"
	"week04/internal/biz"
	"week04/internal/model"

	log "github.com/sirupsen/logrus"
)

type AccountService struct {
	account.UnimplementedAccountServer
	bizCase biz.BizUseCase
}

// 实现grpc接口
func NewAccountService(biz biz.BizUseCase) *AccountService {
	return &AccountService{
		bizCase: biz,
	}
}

func (as *AccountService) CreateUser(ctx context.Context, in *account.CreateUserReq) (*account.CreateUserResp, error) {
	user := &model.User{
		NickName: in.NickName,
		Phone:    in.Phone,
		Age:      in.Age,
		Gender:   in.Gender,
	}
	resp := &account.CreateUserResp{}

	// 这里的 cx 可能是从 grpc 中拿到了一些链路跟踪的信息，然后再将这个 context 传递给 biz 层，因为 biz 可能会调用其他 rpc 服务，这样能够追踪完整的链路情况
	cx := context.TODO()

	err := as.bizCase.CreateUser(cx, user)
	if err != nil {
		log.WithError(err).Errorln("create user failed")
	}
	// 根据不同的业务返回特定的错误码
	resp.Code, resp.Msg = model.MapErr2Code(err)

	return resp, nil
}

func (as *AccountService) DelUser(ctx context.Context, in *account.DelUserReq) (*account.DelUserResp, error) {
	// 因为删除一个用户可能不需要这么多字段，只需要一个userID即可，所以组装对象时使用SimpleUser而不是User结构体。
	user := &model.SimpleUser{
		UserID: in.UserID,
	}
	resp := &account.DelUserResp{}

	// 这里的 cx 可能是从 grpc 中拿到了一些链路跟踪的信息，然后再将这个 context 传递给 biz 层，因为 biz 可能会调用其他 rpc 服务，这样能够追踪完整的链路情况
	cx := context.TODO()

	err := as.bizCase.DelUser(cx, user)
	if err != nil {
		log.WithError(err).Errorln("delete user failed")
	}
	resp.Code, resp.Msg = model.MapErr2Code(err)

	return resp, nil
}
