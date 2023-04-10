package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/validator"
	db "github.com/expert-pancake/service/purchasing/db/transaction"
	"github.com/expert-pancake/service/purchasing/impl"
	"github.com/expert-pancake/service/purchasing/util"
	_ "github.com/lib/pq"
)

type component struct {
	server http.Server
}

func (c *component) New() error {

	config, err := util.LoadConfig()
	if err != nil {
		log.Panic(err)
	}

	dbConnectionSpec := config.Db.GetConnectionSpec()
	conn, err := sql.Open(dbConnectionSpec.DriverName, dbConnectionSpec.DataSourceName)
	if err != nil {
		log.Panic(err)
	}

	purchasingTrx := db.NewPurchasingTrx(conn)

	validator := validator.NewValidator()

	purchasingService := impl.NewPurchasingService(config, validator, purchasingTrx)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Server.Port),
		Handler: c.Routes(purchasingService),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

	return nil
}

func main() {
	log.Println("Starting purchasing service")

	runner := &component{}
	runner.New()

}
