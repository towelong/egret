package usecase

import (
	"context"
	"github.com/towelong/egret/example/internal/domain"
	"go.uber.org/zap"
)

var _ domain.UserUsecase = (*userUsecase)(nil)

type userUsecase struct {
	userRepo domain.UserRepo
	logger   *zap.Logger
}

func NewUserUsecase(userRepo domain.UserRepo, logger *zap.Logger) domain.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (u *userUsecase) GetUserById(ctx context.Context, id int64) (*domain.User, error) {
	return u.userRepo.GetUserById(ctx, id)
}
