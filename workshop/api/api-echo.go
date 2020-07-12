package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

type Users []User

type User struct {
	Id   int64   `json:"id"`  
	Name  string  `json:"name"` 
	Price float64 `json:"price"`
}

func XXX(c echo.Context) error {
	users := Users {
		User{},
		User{},
	}
	return c.JSON(http.StatusOK, users)
}

func YYY(c echo.Context) error {
  u := new(User)
  if err := c.Bind(u); err != nil {
    return err
  }
  // Validation input
  return c.JSON(http.StatusOK, u)
}

func main() {
	e := echo.New()
	e.GET("/", XXX)
	e.POST("/", YYY)
	e.Logger.Fatal(e.Start(":1323"))
}