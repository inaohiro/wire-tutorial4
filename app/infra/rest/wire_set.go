package rest

import (
	"github.com/google/wire"
	"github.com/inaohiro-sandbox/wire-tutorial4/app/infra/rest/controllers"
)

var Set = wire.NewSet(
	NewRouter,
	NewServer,
	controllers.NewUserController,
)
