package vproductapp

import (
	"net/http"

	"github.com/mydomain/see-other/app/sdk/auth"
	"github.com/mydomain/see-other/app/sdk/authclient"
	"github.com/mydomain/see-other/app/sdk/mid"
	"github.com/mydomain/see-other/business/domain/userbus"
	"github.com/mydomain/see-other/business/domain/vproductbus"
	"github.com/mydomain/see-other/foundation/logger"
	"github.com/mydomain/see-other/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log         *logger.Logger
	UserBus     *userbus.Business
	VProductBus *vproductbus.Business
	AuthClient  *authclient.Client
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	authen := mid.Authenticate(cfg.AuthClient)
	ruleAdmin := mid.Authorize(cfg.AuthClient, auth.RuleAdminOnly)

	api := newApp(cfg.VProductBus)

	app.HandlerFunc(http.MethodGet, version, "/vproducts", api.query, authen, ruleAdmin)
}
