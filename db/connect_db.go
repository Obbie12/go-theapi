package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"go-theapi/util"
)

func ConnectPostgresql() *sql.DB {
	var err error
	config, _ := util.LoadConfig(".")

	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}

	return db
}
