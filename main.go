package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/alterra-kelompok-1/menu-service/database"
	"github.com/alterra-kelompok-1/menu-service/database/migration"
	"github.com/alterra-kelompok-1/menu-service/database/seeder"
	"github.com/alterra-kelompok-1/menu-service/internal/factory"
	"github.com/alterra-kelompok-1/menu-service/internal/http"
	"github.com/alterra-kelompok-1/menu-service/internal/middleware"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// load env configuration
func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {

	database.CreateConnection()

	var m string // for check migration

	flag.StringVar(
		&m,
		"migrate",
		"run",
		`this argument for check if menu want to migrate table, rollback table, or status migration

to use this flag:
	use -migrate=migrate for migrate table
	use -migrate=rollback for rollback table
	use -migrate=status for get status migration`,
	)
	flag.Parse()

	if m == "migrate" {
		migration.Migrate()
		return
	} else if m == "rollback" {
		migration.Rollback()
		return
	} else if m == "status" {
		migration.Status()
		return
	}

	// TEST
	fmt.Println(database.GetConnection())

	seeder.Seed()

	f := factory.NewFactory()
	e := echo.New()
	middleware.Init(e)
	http.NewHttp(e, f)

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))

}
