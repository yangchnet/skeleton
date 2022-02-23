package biz

import "github.com/google/wire"

// ProviderSet provided NewEchoCase.
var ProviderSet = wire.NewSet(NewEchoCase)
