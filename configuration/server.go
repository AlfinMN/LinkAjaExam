package configuration

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func CreateRouter() *echo.Echo {
	e := echo.New()

	return e
}
func RunServer(e *echo.Echo, serverHost, serverPort string) {

	fmt.Println("succes connect to port : " + serverPort)
	err := http.ListenAndServe(serverHost+":"+serverPort, e)
	if err != nil {
		log.Fatal(err)
	}
}
