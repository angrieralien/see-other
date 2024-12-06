package homeapp

import (
	"net/http"

	"github.com/angrieralien/seeother/app/sdk/auth"
	"github.com/angrieralien/seeother/app/sdk/authclient"
	"github.com/angrieralien/seeother/app/sdk/mid"
	"github.com/angrieralien/seeother/business/domain/homebus"
	"github.com/angrieralien/seeother/business/domain/userbus"
	"github.com/angrieralien/seeother/foundation/logger"
	"github.com/angrieralien/seeother/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log        *logger.Logger
	UserBus    *userbus.Business
	HomeBus    *homebus.Business
	AuthClient *authclient.Client
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	authen := mid.Authenticate(cfg.AuthClient)
	ruleAny := mid.Authorize(cfg.AuthClient, auth.RuleAny)
	ruleUserOnly := mid.Authorize(cfg.AuthClient, auth.RuleUserOnly)
	ruleAuthorizeHome := mid.AuthorizeHome(cfg.AuthClient, cfg.HomeBus)

	api := newApp(cfg.HomeBus)

	app.HandlerFunc(http.MethodGet, version, "/homes", api.query, authen, ruleAny)
	app.HandlerFunc(http.MethodGet, version, "/homes/{home_id}", api.queryByID, authen, ruleAuthorizeHome)
	app.HandlerFunc(http.MethodPost, version, "/homes", api.create, authen, ruleUserOnly)
	app.HandlerFunc(http.MethodPut, version, "/homes/{home_id}", api.update, authen, ruleAuthorizeHome)
	app.HandlerFunc(http.MethodDelete, version, "/homes/{home_id}", api.delete, authen, ruleAuthorizeHome)
}
