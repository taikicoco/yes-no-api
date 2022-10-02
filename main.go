package main

import (
	"api/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var err error

func main() {
  e := echo.New()

  // middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  if err = handler.Routes(e); err != nil {
    panic(err)
  }

  // start server
  e.Logger.Fatal(e.Start(":1333"))
}


