// Package vproductapp maintains the app layer api for the vproduct domain.
package vproductapp

import (
	"context"
	"net/http"

	"github.com/angrieralien/seeother/app/sdk/errs"
	"github.com/angrieralien/seeother/app/sdk/query"
	"github.com/angrieralien/seeother/business/domain/vproductbus"
	"github.com/angrieralien/seeother/business/sdk/order"
	"github.com/angrieralien/seeother/business/sdk/page"
	"github.com/angrieralien/seeother/foundation/web"
)

type app struct {
	vproductBus *vproductbus.Business
}

func newApp(vproductBus *vproductbus.Business) *app {
	return &app{
		vproductBus: vproductBus,
	}
}

func (a *app) query(ctx context.Context, r *http.Request) web.Encoder {
	qp := parseQueryParams(r)

	page, err := page.Parse(qp.Page, qp.Rows)
	if err != nil {
		return errs.NewFieldsError("page", err)
	}

	filter, err := parseFilter(qp)
	if err != nil {
		return err.(errs.FieldErrors)
	}

	orderBy, err := order.Parse(orderByFields, qp.OrderBy, vproductbus.DefaultOrderBy)
	if err != nil {
		return errs.NewFieldsError("order", err)
	}

	prds, err := a.vproductBus.Query(ctx, filter, orderBy, page)
	if err != nil {
		return errs.Newf(errs.Internal, "query: %s", err)
	}

	total, err := a.vproductBus.Count(ctx, filter)
	if err != nil {
		return errs.Newf(errs.Internal, "count: %s", err)
	}

	return query.NewResult(toAppProducts(prds), total, page)
}
