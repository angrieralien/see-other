package dbtest

import (
	"time"

	"github.com/angrieralien/seeother/business/domain/homebus"
	"github.com/angrieralien/seeother/business/domain/homebus/stores/homedb"
	"github.com/angrieralien/seeother/business/domain/productbus"
	"github.com/angrieralien/seeother/business/domain/productbus/stores/productdb"
	"github.com/angrieralien/seeother/business/domain/userbus"
	"github.com/angrieralien/seeother/business/domain/userbus/stores/usercache"
	"github.com/angrieralien/seeother/business/domain/userbus/stores/userdb"
	"github.com/angrieralien/seeother/business/domain/vproductbus"
	"github.com/angrieralien/seeother/business/domain/vproductbus/stores/vproductdb"
	"github.com/angrieralien/seeother/business/sdk/delegate"
	"github.com/angrieralien/seeother/foundation/logger"
	"github.com/jmoiron/sqlx"
)

// BusDomain represents all the business domain apis needed for testing.
type BusDomain struct {
	Delegate *delegate.Delegate
	Home     *homebus.Business
	Product  *productbus.Business
	User     *userbus.Business
	VProduct *vproductbus.Business
}

func newBusDomains(log *logger.Logger, db *sqlx.DB) BusDomain {
	delegate := delegate.New(log)
	userBus := userbus.NewBusiness(log, delegate, usercache.NewStore(log, userdb.NewStore(log, db), time.Hour))
	productBus := productbus.NewBusiness(log, userBus, delegate, productdb.NewStore(log, db))
	homeBus := homebus.NewBusiness(log, userBus, delegate, homedb.NewStore(log, db))
	vproductBus := vproductbus.NewBusiness(vproductdb.NewStore(log, db))

	return BusDomain{
		Delegate: delegate,
		Home:     homeBus,
		Product:  productBus,
		User:     userBus,
		VProduct: vproductBus,
	}
}
