package main

import (
	"api-devbook/src/config"
	"api-devbook/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()
	r := router.Generate()
	log.Println("Rodando")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
