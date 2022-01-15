package usecase

import "github.com/google/wire"

// ProvideSet  is usecase providers.
var ProvideSet = wire.NewSet(NewUserUsecase)
