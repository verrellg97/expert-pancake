package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/validator"
	db "github.com/expert-pancake/service/inventory/db/transaction"
	"github.com/expert-pancake/service/inventory/impl"
	"github.com/expert-pancake/service/inventory/util"
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

	inventoryTrx := db.NewInventoryTrx(conn)

	validator := validator.NewValidator()

	inventoryService := impl.NewInventoryService(config, validator, inventoryTrx)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Server.Port),
		Handler: c.Routes(inventoryService),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

	return nil
}

func main() {
	log.Println("Starting inventory service")

	runner := &component{}
	runner.New()

}
