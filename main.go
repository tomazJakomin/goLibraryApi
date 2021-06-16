package main

import (
	"fmt"
	"github.com/tomazJakomin/go-base-app/api/routing"
	"github.com/tomazJakomin/go-base-app/internal/config"
	"github.com/tomazJakomin/go-base-app/internal/database"
	"net/http"
)

func main() {
	config := config.NewConfig()
	db := database.NewDb(*config)

	if db.Error != nil {
		panic(db.Error)
	}

	router := routing.StartRouter(db)

	fmt.Printf("Starting server %s:%s...", config.Server.Host, config.Server.Port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", config.Server.Port), router)
	if err != nil {
		fmt.Printf("Error on server %s", err)
	}
}
