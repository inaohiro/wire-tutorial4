//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/inaohiro-sandbox/wire-tutorial4/app/config"
	"github.com/inaohiro-sandbox/wire-tutorial4/app/domain"
	"github.com/inaohiro-sandbox/wire-tutorial4/app/infra/database/mysql"
	"github.com/inaohiro-sandbox/wire-tutorial4/app/infra/rest"
	"github.com/inaohiro-sandbox/wire-tutorial4/app/interfaces/database"
	"github.com/inaohiro-sandbox/wire-tutorial4/app/interfaces/inmemorydb"
)

func initializeRealServer() (*rest.Server, func(), error) {
	wire.Build(
		domain.Set,
		mysql.Set,
		config.Set,
		database.Set,
		rest.Set,
	)
	return nil, nil, nil
}

func initializeServer() (*rest.Server, func(), error) {
	wire.Build(
		domain.Set,
		config.Set,
		inmemorydb.NewUserRepository,
		wire.Bind(new(domain.UserRepository), new(*inmemorydb.InMemoryUserRepository)),
		rest.Set,
	)
	return nil, nil, nil
}
