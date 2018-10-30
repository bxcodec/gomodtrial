package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/bxcodec/gomodmultiplies/v2"
	"github.com/labstack/echo"
)

func main() {
	a, b := int64(10), int64(23)
	res := gomodmultiplies.Multiply(a, b)
	fmt.Println(res)

	c, d := 10, 20
	r := gomodmultiplies.Multiply(c, d)
	fmt.Println(r)

	e := echo.New()
	e.GET("/multiply", multiplyHandler)
	e.Start(":9090")
}

func multiplyHandler(c echo.Context) error {
	params := c.QueryParam("params")
	arrParams := strings.Split(params, ",")
	arrParamInt64 := []interface{}{}
	for _, p := range arrParams {
		param, err := strconv.ParseInt(p, 10, 64)
		if err != nil {
			return err
		}
		arrParamInt64 = append(arrParamInt64, param)
	}

	res := gomodmultiplies.Multiply(arrParamInt64...)
	return c.JSON(http.StatusOK, res)
}
