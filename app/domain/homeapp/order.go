package homeapp

import (
	"github.com/mydomain/see-other/business/domain/homebus"
)

var orderByFields = map[string]string{
	"home_id": homebus.OrderByID,
	"type":    homebus.OrderByType,
	"user_id": homebus.OrderByUserID,
}
