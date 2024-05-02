package main

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/henrique77/product-api/infra/db"
	"github.com/henrique77/product-api/src/handlers"
	"github.com/henrique77/product-api/src/modules/product/repositories"
	service "github.com/henrique77/product-api/src/modules/product/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := db.DB.Open()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	r := repositories.New(db)
	s := service.New(r)
	h := handlers.New(s)

	handlers.SetApi(e, h)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8888"
	}

	e.Logger.Fatal(e.Start(":" + PORT))

}
