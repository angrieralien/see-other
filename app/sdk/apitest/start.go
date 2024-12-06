package apitest

import (
	"net/http/httptest"
	"testing"

	authbuild "github.com/angrieralien/seeother/api/services/auth/build/all"
	salesbuild "github.com/angrieralien/seeother/api/services/sales/build/all"
	"github.com/angrieralien/seeother/app/sdk/auth"
	"github.com/angrieralien/seeother/app/sdk/authclient"
	"github.com/angrieralien/seeother/app/sdk/mux"
	"github.com/angrieralien/seeother/business/sdk/dbtest"
)

// New initialized the system to run a test.
func New(t *testing.T, testName string) *Test {
	db := dbtest.New(t, testName)

	// -------------------------------------------------------------------------

	auth, err := auth.New(auth.Config{
		Log:       db.Log,
		DB:        db.DB,
		KeyLookup: &KeyStore{},
	})
	if err != nil {
		t.Fatal(err)
	}

	// -------------------------------------------------------------------------

	server := httptest.NewServer(mux.WebAPI(mux.Config{
		Log: db.Log,
		DB:  db.DB,
		AuthConfig: mux.AuthConfig{
			Auth: auth,
		},
	}, authbuild.Routes()))

	authClient := authclient.New(db.Log, server.URL)

	// -------------------------------------------------------------------------

	mux := mux.WebAPI(mux.Config{
		Log: db.Log,
		DB:  db.DB,
		SalesConfig: mux.SalesConfig{
			AuthClient: authClient,
		},
	}, salesbuild.Routes())

	return &Test{
		DB:   db,
		Auth: auth,
		mux:  mux,
	}
}
