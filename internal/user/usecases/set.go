package usecases

import "github.com/google/wire"

var Set = wire.NewSet(NewUserUsecase)
