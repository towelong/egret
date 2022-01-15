package repo

import "github.com/google/wire"

// ProvideSet  is repo providers.
var ProvideSet = wire.NewSet(NewUserRepo)
