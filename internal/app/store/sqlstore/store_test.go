package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "user=leo password=123 host=localhost dbname=restapi_test sslmode=disable"
	}
	os.Exit(m.Run())
}