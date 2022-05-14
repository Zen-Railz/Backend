//go:build wireinject
// +build wireinject

package system

import (
	"database/sql"
	"zenrailz/log"
	"zenrailz/nexus"
	"zenrailz/repository"
	configRepo "zenrailz/repository/configuration"
	databaseRepo "zenrailz/repository/database"
	railwayRepo "zenrailz/repository/railway"
	"zenrailz/service"
	"zenrailz/service/health"
	"zenrailz/service/railway"

	"github.com/google/wire"
)

var databaseRepositorySet = wire.NewSet(
	databaseRepo.NewRepository,
	wire.Bind(new(repository.Database), new(*databaseRepo.Repository)),
)

var configurationRepositorySet = wire.NewSet(
	configRepo.NewRepository,
	wire.Bind(new(repository.Configuration), new(*configRepo.Repository)),
)

var railwayRepositorySet = wire.NewSet(
	railwayRepo.NewRepository,
	wire.Bind(new(repository.Railway), new(*railwayRepo.Repository)),
)

var healthServiceSet = wire.NewSet(
	health.NewService,
	wire.Bind(new(service.Health), new(*health.Service)),
)

var railwayServiceSet = wire.NewSet(
	railway.NewService,
	wire.Bind(new(service.Railway), new(*railway.Service)),
)

func InitialiseNexus(logger log.Logger, db *sql.DB) *nexus.Store {
	panic(
		wire.Build(
			databaseRepositorySet,
			configurationRepositorySet,
			railwayRepositorySet,

			healthServiceSet,
			railwayServiceSet,

			nexus.NewStore,
		),
	)
}
