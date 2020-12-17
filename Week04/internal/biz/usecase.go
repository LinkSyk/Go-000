package biz

import (
	"context"
	"week04/internal/model"
	"week04/internal/repo"
)

type BizUseCase interface {
	CreateUser(ctx context.Context, user *model.User) error
	DelUser(ctx context.Context, user *model.SimpleUser) error
}

type BizUseCaseImpl struct {
	mysqlRepo repo.DBRepository
}

func NewBizImpl(mysql repo.DBRepository) BizUseCase {
	return &BizUseCaseImpl{
		mysqlRepo: mysql,
	}
}

func (b *BizUseCaseImpl) CreateUser(ctx context.Context, user *model.User) error {
	return b.mysqlRepo.CreateUser(user)
}

func (b *BizUseCaseImpl) DelUser(ctx context.Context, user *model.SimpleUser) error {
	return b.mysqlRepo.DelUser(user)
}
