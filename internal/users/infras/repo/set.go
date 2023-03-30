package repo

import "github.com/google/wire"

var RepositorySet = wire.NewSet(NewUserRepository, NewUserCredentialsRepository)
