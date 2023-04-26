package db_test

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/wjhdec/echo-ext/pkg/config"
	"github.com/wjhdec/echo-ext/pkg/db"
	"testing"
)

func TestName(t *testing.T) {
	cfg, err := config.New("../../configs")
	if err != nil {
		panic(err)
	}
	mdb, err := db.New(cfg).GetByName("local")
	if err != nil {
		panic(err)
	}
	defer mdb.Close()
	row := mdb.QueryRow("select version()")
	if err != nil {
		panic(err)
	}
	var version string
	if err := row.Scan(&version); err != nil {
		panic(err)
	}
	fmt.Println(version)
}
