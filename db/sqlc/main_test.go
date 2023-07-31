package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/techschool/simplebank/util"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Can not load config: ", err)
	}
	testDB, err = sql.Open(config.DBdriver, config.DBsource)
	if err != nil {
		log.Fatal("cannot connect database:", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
