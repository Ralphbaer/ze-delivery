//+build wireinject
//golint:ignore

package gen

import (
	"net/http"
	"sync"

	"github.com/Ralphbaer/ze-delivery/common"
	"github.com/Ralphbaer/ze-delivery/partner/app"
	h "github.com/Ralphbaer/ze-delivery/partner/handler"
	r "github.com/Ralphbaer/ze-delivery/partner/repository"
	uc "github.com/Ralphbaer/ze-delivery/partner/usecase"
	"github.com/google/wire"
	"github.com/gorilla/mux"
)

var onceConfig sync.Once

func setupMongoConnection(cfg *app.Config) *common.MongoConnection {
	return &common.MongoConnection{
		ConnectionString: cfg.MongoConnectionString,
		Verbose:          true,
	}
}

var applicationSet = wire.NewSet(
	common.InitLocalEnvConfig,
	setupMongoConnection,
	app.NewConfig,
	app.NewRouter,
	app.NewServer,
	r.NewPartnerMongoRepository,
	wire.Struct(new(uc.PartnerUseCase), "*"),
	wire.Struct(new(h.PartnerHandler), "*"),
	wire.Bind(new(r.PartnerRepository), new(*r.PartnerMongoRepository)),
	wire.Bind(new(http.Handler), new(*mux.Router)),
)

// InitializeApp setup the dependencies and returns a new *app.App instance
func InitializeApp() *app.App {
	wire.Build(
		applicationSet,
		wire.Struct(new(app.App), "*"),
	)
	return nil
}
