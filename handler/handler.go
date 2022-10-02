package handler

import (
	"api/data"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"math/rand"

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

	res := resp[random()].Answer
	return e.JSON(http.StatusOK, res)
}

func random() int {
	v := rand.Intn(100)

	if v < 45 {
		return 0
	} else if v < 90 {
		return 1
	} else {
		return 2
	}
}