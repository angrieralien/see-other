// Package all binds all the routes into the specified app.
package all

import (
	"time"

	"github.com/mydomain/see-other/app/domain/checkapp"
	"github.com/mydomain/see-other/app/domain/homeapp"
	"github.com/mydomain/see-other/app/domain/productapp"
	"github.com/mydomain/see-other/app/domain/rawapp"
	"github.com/mydomain/see-other/app/domain/tranapp"
	"github.com/mydomain/see-other/app/domain/userapp"
	"github.com/mydomain/see-other/app/domain/vproductapp"
	"github.com/mydomain/see-other/app/sdk/mux"
	"github.com/mydomain/see-other/business/domain/homebus"
	"github.com/mydomain/see-other/business/domain/homebus/stores/homedb"
	"github.com/mydomain/see-other/business/domain/productbus"
	"github.com/mydomain/see-other/business/domain/productbus/stores/productdb"
	"github.com/mydomain/see-other/business/domain/userbus"
	"github.com/mydomain/see-other/business/domain/userbus/stores/usercache"
	"github.com/mydomain/see-other/business/domain/userbus/stores/userdb"
	"github.com/mydomain/see-other/business/domain/vproductbus"
	"github.com/mydomain/see-other/business/domain/vproductbus/stores/vproductdb"
	"github.com/mydomain/see-other/business/sdk/delegate"
	"github.com/mydomain/see-other/foundation/web"
)

// Routes constructs the add value which provides the implementation of
// of RouteAdder for specifying what routes to bind to this instance.
func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouterAdder interface.
func (add) Add(app *web.App, cfg mux.Config) {

	// Construct the business domain packages we need here so we are using the
	// sames instances for the different set of domain apis.
	delegate := delegate.New(cfg.Log)
	userBus := userbus.NewBusiness(cfg.Log, delegate, usercache.NewStore(cfg.Log, userdb.NewStore(cfg.Log, cfg.DB), time.Minute))
	productBus := productbus.NewBusiness(cfg.Log, userBus, delegate, productdb.NewStore(cfg.Log, cfg.DB))
	homeBus := homebus.NewBusiness(cfg.Log, userBus, delegate, homedb.NewStore(cfg.Log, cfg.DB))
	vproductBus := vproductbus.NewBusiness(vproductdb.NewStore(cfg.Log, cfg.DB))

	checkapp.Routes(app, checkapp.Config{
		Build: cfg.Build,
		Log:   cfg.Log,
		DB:    cfg.DB,
	})

	homeapp.Routes(app, homeapp.Config{
		Log:        cfg.Log,
		UserBus:    userBus,
		HomeBus:    homeBus,
		AuthClient: cfg.AuthClient,
	})

	productapp.Routes(app, productapp.Config{
		Log:        cfg.Log,
		UserBus:    userBus,
		ProductBus: productBus,
		AuthClient: cfg.AuthClient,
	})

	rawapp.Routes(app)

	tranapp.Routes(app, tranapp.Config{
		Log:        cfg.Log,
		DB:         cfg.DB,
		UserBus:    userBus,
		ProductBus: productBus,
		AuthClient: cfg.AuthClient,
	})

	userapp.Routes(app, userapp.Config{
		Log:        cfg.Log,
		UserBus:    userBus,
		AuthClient: cfg.AuthClient,
	})

	vproductapp.Routes(app, vproductapp.Config{
		Log:         cfg.Log,
		UserBus:     userBus,
		VProductBus: vproductBus,
		AuthClient:  cfg.AuthClient,
	})
}
