package handler

import (
	"api/data"
	_ "api/data"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) error {

	e.GET("/",hello)
	e.GET("/get",get)

	return nil
}

func hello(e echo.Context) error {
	return e.String(http.StatusOK,"heelo")
}

func get(e echo.Context) error {

	var resp data.Answer
	f, err := os.Open("./data/data.json")
	
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()


	err = json.Unmarshal(content, &resp)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp[1])
	return e.JSON(http.StatusOK, resp)
}

