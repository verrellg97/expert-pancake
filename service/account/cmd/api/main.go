package main

import (
	"database/sql"
	"fmt"
	"github.com/calvinkmts/expert-pancake/engine/validator"
	db "github.com/expert-pancake/service/account/db/transaction"
	"github.com/expert-pancake/service/account/impl"
	"github.com/expert-pancake/service/account/token"
	"github.com/expert-pancake/service/account/util"
	_ "github.com/lib/pq"
	"log"
	"net/http"
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

	accountTrx := db.NewAccountTrx(conn)

	validator := validator.NewValidator()

	tokenMaker, err := token.NewPasetoMaker(config.Token.SymmetricKey)
	if err != nil {
		log.Panic(err)
	}

	accountService := impl.NewAccountService(config, validator, accountTrx, tokenMaker)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Server.Port),
		Handler: c.Routes(accountService),
	}

	err = server.ListenAndServe()
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
