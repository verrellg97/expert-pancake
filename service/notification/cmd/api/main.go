package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/validator"
	db "github.com/expert-pancake/service/notification/db/transaction"
	"github.com/expert-pancake/service/notification/impl"
	"github.com/expert-pancake/service/notification/util"
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

	notificationTrx := db.NewNotificationTrx(conn)

	validator := validator.NewValidator()

	notificationService := impl.NewNotificationService(config, validator, notificationTrx)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Server.Port),
		Handler: c.Routes(notificationService),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

	return nil
}

func main() {
	log.Println("Starting notification service")

	runner := &component{}
	runner.New()

}
