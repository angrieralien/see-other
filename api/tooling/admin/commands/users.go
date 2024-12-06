package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/mydomain/see-other/business/domain/userbus"
	"github.com/mydomain/see-other/business/domain/userbus/stores/userdb"
	"github.com/mydomain/see-other/business/sdk/page"
	"github.com/mydomain/see-other/business/sdk/sqldb"
	"github.com/mydomain/see-other/foundation/logger"
)

// Users retrieves all users from the database.
func Users(log *logger.Logger, cfg sqldb.Config, pageNumber string, rowsPerPage string) error {
	db, err := sqldb.Open(cfg)
	if err != nil {
		return fmt.Errorf("connect database: %w", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userBus := userbus.NewBusiness(log, nil, userdb.NewStore(log, db))

	page, err := page.Parse(pageNumber, rowsPerPage)
	if err != nil {
		return fmt.Errorf("parsing page information: %w", err)
	}

	users, err := userBus.Query(ctx, userbus.QueryFilter{}, userbus.DefaultOrderBy, page)
	if err != nil {
		return fmt.Errorf("retrieve users: %w", err)
	}

	return json.NewEncoder(os.Stdout).Encode(users)
}
