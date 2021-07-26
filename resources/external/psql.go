package external

import (
	"github.com/ravielze/Labpro-BE/config"
	"github.com/ravielze/oculi/logs"
	"github.com/ravielze/oculi/persistent/sql"
	"github.com/ravielze/oculi/persistent/sql/postgre"
	sqlv2 "github.com/ravielze/oculi/persistent/sql/v2"
)

type (
	DBManager struct {
		Install func()
		Reset   func()
	}
)

func NewPostgreSQL(config *config.Env, log logs.Logger) (sql.API, error) {
	return sqlv2.NewClient(
		postgre.New(sql.ConnectionInfo{
			Address:  config.DatabaseAddress,
			Username: config.DatabaseUsername,
			Password: config.DatabasePassword,
			DbName:   config.DatabaseName,
		}), false,
		sql.WithMaxIdleConnection(config.DatabaseMaxIdleConnection),
		sql.WithMaxOpenConnection(config.DatabaseMaxOpenConnection),
		sql.WithConnMaxLifetime(config.DatabaseConnMaxLifetime),
		sql.WithLogMode(config.DatabaseLogMode),
		sql.WithLogger(log))
}

func NewDBManager(api sql.API) *DBManager {
	api.RegisterObject()
	i, r := api.ObjectFunction(nil, nil)
	i()
	return &DBManager{
		Install: i,
		Reset:   r,
	}
}
