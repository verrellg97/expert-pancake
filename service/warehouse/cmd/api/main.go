package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/validator"
	db "github.com/expert-pancake/service/warehouse/db/transaction"
	"github.com/expert-pancake/service/warehouse/impl"
	"github.com/expert-pancake/service/warehouse/util"
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

	warehouseTrx := db.NewWarehouseTrx(conn)

	validator := validator.NewValidator()

	warehouseService := impl.NewWarehouseService(config, validator, warehouseTrx)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Server.Port),
		Handler: c.Routes(warehouseService),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

	return nil
}

func main() {
	log.Println("Starting warehouse service")

	runner := &component{}
	runner.New()

}
