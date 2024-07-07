package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/muizu555/investment/src/repository"
	"github.com/muizu555/investment/src/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	router.SetupRoutes(e)

	err = repository.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer repository.GetDB().Close()

	fmt.Println("Server is running at :8080")
	err = e.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
