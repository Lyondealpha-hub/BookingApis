package database

import (
	"fmt"
	"log"

	"andurar.api/config"
	"andurar.api/ent"
)

func postgresConnector(conf *config.DBConfig) *ent.Client {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		conf.Host,
		conf.Port,
		conf.Username,
		conf.Name,
		conf.Password,
		conf.SSL,
	)
	client, err := ent.Open(conf.Driver, dsn)
	if err != nil {
		log.Fatalf("failed connecting to postgres: %v", err)
	}
	return client
}
