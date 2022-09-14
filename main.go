package main

import (
	_"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
  
  e := echo.New()

  // ミドルウェア
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // ルート
  e.GET("/", hello) 
  e.GET("/get/:id", get)
  e.GET("/json",getjson)
  e.POST("getname",getname)

  // サーバーをポート番号1323で起動
  e.Logger.Fatal(e.Start(":1323"))
  
}

// ハンドラーを定義
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}

func get(c echo.Context) error {
  id := c.Param("id")
  return c.String(http.StatusOK, id)
}

func getjson(c echo.Context) error {
  return c.JSON(http.StatusOK, map[string]string{"hello": "world"})
}

func getname(c echo.Context) error {
  param := new(HelloParam)
    if err := c.Bind(&param); err != nil {
        return err
    }
    return c.JSON(http.StatusOK, map[string]interface{}{"hello": param.GreetingTo})
}

type User struct {
  Name  string `json:"name"`
}

type HelloParam struct {
  GreetingTo string `json:"greetingto"`
}