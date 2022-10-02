package handler

import (
	"api/data"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) error {

	e.GET("/",get_json)

	return nil
}

func get_json(e echo.Context) error {

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

