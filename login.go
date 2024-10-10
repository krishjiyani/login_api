package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Request struct {
	Email    string `  json:"email" `
	Password string ` json:"password" `
}

/*
user: {
key  : value
email: krish@123

}
*/
type Response struct {
	Message string ` json: "message" `
}

func SimpleHandler(ctx echo.Context) error {

	var req Request
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, Response{Message: "empty email,please provide email and password"})

	}

	if req.Email == "" || req.Password == "" {
		return ctx.JSON(http.StatusBadRequest, Response{Message: "empty email,please provide email and password"})
	}
	message := "login successfullll!! "

	return ctx.JSON(http.StatusOK, Response{Message: message})

}
func main() {
	e := echo.New()

	e.POST("/login", SimpleHandler)
	e.Start(":8080")
}
