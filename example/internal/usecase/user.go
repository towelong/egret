package usecase

import (
	"context"
	"github.com/towelong/egret/example/internal/domain"
	"go.uber.org/zap"
)

var _ domain.UserUsecase = (*UserUsecase)(nil)

type UserUsecase struct {
	userRepo domain.UserRepo
	logger   *zap.Logger
}

func NewUserUsecase(userRepo domain.UserRepo, logger *zap.Logger) domain.UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (u *UserUsecase) GetUserById(ctx context.Context, id int64) (*domain.User, error) {
	return u.userRepo.GetUserById(ctx, id)
}
