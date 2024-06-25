package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"andurar.api/config"
	"andurar.api/ent"
	"andurar.api/ent/migrate"
)

type Adapter struct {
	DB *ent.Client
}

func NewDB() (*ent.Client, error) {
	conf := config.DB()
	var dsn string
	switch conf.Driver {
	case "mysql":
		fmt.Sprint("%s:%s@tcp(%s:%s)/%s?parseTime=True",
			conf.Username,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.Name,
		)

		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=True",
			conf.Username,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.Name,
		)
	case "postgres":
		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			conf.Host,
			conf.Port,
			conf.Username,
			conf.Password,
			conf.Name,
			conf.SSL,
		)
	default:
		return nil, fmt.Errorf("unsupported driver: %s", conf.Driver)
	}

	log.Println("driver and dssn: ", conf.Driver, dsn)

	client, err := ent.Open(conf.Driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to %s: %w", conf.Driver, err)
	}

	return client, nil
}

func MigrateDB(db *ent.Client) error {
	// ctx := context.Background()
	log.Println("Starting database migration...")

	if err := db.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		return fmt.Errorf("failed creating schema resources: %v", err)
	}

	log.Println("Database migration completed successfully.")
	return nil
}
func Connect() *sql.DB {
	conf := config.DB()
	switch conf.Driver {
	case "postgres":
		psDSN := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			conf.Host,
			conf.Port,
			conf.Username,
			conf.Name,
			conf.Password,
			conf.SSL,
		)
		db, err := sql.Open(conf.Driver, psDSN)
		if err != nil {
			log.Panicln(err.Error())
		}
		return db
	default:
		mysqlDSN := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			conf.Username,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.Name,
		)
		db, err := sql.Open(conf.Driver, mysqlDSN)
		if err != nil {
			log.Panicln(err.Error())
		}
		return db
	}
}
func ConnectServer() *sql.DB {
	conf := config.DB()
	switch conf.Driver {
	case "postgres":
		psDSN := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s sslmode=%s",
			conf.Host,
			conf.Port,
			conf.Username,
			conf.Password,
			conf.SSL,
		)
		db, err := sql.Open(conf.Driver, psDSN)
		if err != nil {
			log.Panicln(err.Error())
		}
		return db
	default:
		mysqlDSN := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/?multiStatements=true",
			conf.Username,
			conf.Password,
			conf.Host,
			conf.Port,
		)
		db, err := sql.Open(conf.Driver, mysqlDSN)
		if err != nil {
			log.Panicln(err.Error())
		}

		return db
	}
}
func CreateDB(conn *sql.DB, database string) error {
	_, err := conn.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s;CREATE DATABASE IF NOT EXISTS %s;", database, database))
	if err != nil {
		return err
	}
	return nil
}
