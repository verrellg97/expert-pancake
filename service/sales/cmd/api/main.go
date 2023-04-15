package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/validator"
	db "github.com/expert-pancake/service/sales/db/transaction"
	"github.com/expert-pancake/service/sales/impl"
	"github.com/expert-pancake/service/sales/util"
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

	salesTrx := db.NewSalesTrx(conn)

	validator := validator.NewValidator()

	salesService := impl.NewSalesService(config, validator, salesTrx)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Server.Port),
		Handler: c.Routes(salesService),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

	return nil
}

func main() {
	log.Println("Starting sales service")

	runner := &component{}
	runner.New()

}
