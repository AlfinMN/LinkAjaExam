package master

import (
	"database/sql"
	"testing"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func DBTest() *sql.DB {
	db, _ := sql.Open("sqlite3", ":memory:")
	return db
}

func RouterTest() *echo.Echo {
	e := echo.New()

	return e
}

func TestInit(t *testing.T) {
	db := DBTest()
	e := RouterTest()
	InitData(e, db)
}
