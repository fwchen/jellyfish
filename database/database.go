package database

import (
	"database/sql"
	configs "github.com/fwchen/jellyfish/config"
	"github.com/juju/errors"
	_ "github.com/lib/pq"
	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/pq"
)

type AppDataSource struct {
	RDS *sql.DB
}

func connectDatabase(rdsConfig configs.RDSConfig) (*sql.DB, error) {
	db, err := apmsql.Open("postgres", rdsConfig.DatabaseUrl)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return db, nil
}

func GetDatabase(dataSourceConfig configs.DataSourceConfig) (*AppDataSource, error) {
	db, err := connectDatabase(dataSourceConfig.RDS)
	if err != nil {
		return nil, err
	}
	return &AppDataSource{RDS: db}, nil
}
