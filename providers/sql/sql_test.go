package sql

import (
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/scoville/validations"
	"gopkg.in/khaiql/dbcleaner.v2"
	dbengine "gopkg.in/khaiql/dbcleaner.v2/engine"
)

var cleaner = dbcleaner.New()

var testDBDriver = "postgres"
var testDBURL = "host=localhost dbname=encourage_contents_test sslmode=disable"

func TestMain(m *testing.M) {
	if dbURL := os.Getenv("TEST_DB_URL"); dbURL != "" {
		testDBURL = dbURL
	}

	f, err := NewStorage(testDBDriver, testDBURL)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	f.DropTables()
	f.Migrate()

	postgresql := dbengine.NewPostgresEngine(testDBURL)
	cleaner.SetEngine(postgresql)

	retCode := m.Run()
	os.Exit(retCode)
}

func TestRunSuites(t *testing.T) {
	db, err := gorm.Open(testDBDriver, testDBURL)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer db.Close()
	validations.RegisterCallbacks(db)
}
