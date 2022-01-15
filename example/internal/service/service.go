package service

import (
	"github.com/google/wire"
	"github.com/towelong/egret/example/internal/domain"
	"go.uber.org/zap"
)

// ProvideSet  is service providers.
var ProvideSet = wire.NewSet(NewShopInterface)

type ShopInterface struct {
	userUsecase domain.UserUsecase
	logger      *zap.Logger
}

func NewShopInterface(usecase domain.UserUsecase, logger *zap.Logger) *ShopInterface {
	return &ShopInterface{
		userUsecase: usecase,
		logger:      logger,
	}
}
