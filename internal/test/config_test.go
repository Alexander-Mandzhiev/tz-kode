package test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_TEST_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost port=5436 user=postgres password=root dbname=notes_test sslmode=disable"
	}
	os.Exit(m.Run())
}
