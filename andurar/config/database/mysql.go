package database

import (
	"fmt"
	"log"

	"andurar.api/config"
	"andurar.api/ent"
)

func mysqlConnector(conf *config.DBConfig) *ent.Client {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=True",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Name,
	)
	client, err := ent.Open(conf.Driver, dsn)
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}

	return client
}
