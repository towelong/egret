package repo

import (
	"context"
	"github.com/towelong/egret/example/internal/domain"
	"go.uber.org/zap"
)

var _ domain.UserRepo = (*userRepo)(nil)

type userRepo struct {
	logger *zap.Logger
}

func NewUserRepo(logger *zap.Logger) domain.UserRepo {
	return &userRepo{
		logger: logger,
	}
}

func (u *userRepo) GetUserById(ctx context.Context, id int64) (*domain.User, error) {
	u.logger.Info("Repo被用到了")
	return &domain.User{
		ID:   1,
		Name: "Welong",
		Age:  20,
	}, nil
}
