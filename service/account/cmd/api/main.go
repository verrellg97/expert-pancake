package main

import (
	"fmt"
	"github.com/expert-pancake/service/account/impl"
	"log"
	"net/http"
)

type component struct {
	server http.Server
}

func (c *component) New() error {
	accountService := impl.NewAccountService()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", "3000"),
		Handler: c.Routes(accountService),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

	return nil
}

func main() {
	log.Println("Starting account service")

	runner := &component{}
	runner.New()

}
