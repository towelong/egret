package server

import "github.com/google/wire"

// ProvideSet  is http providers.
var ProvideSet = wire.NewSet(NewHttpServer)
