package db

import (
	"fristTry/config"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

var Db *bun.DB
var err error

func CreateDatabase(configurations config.Configurations) *bun.DB {
	config, err := pgx.ParseConfig(configurations.Db.Connection_string)
	if err != nil {
		panic(err)
	}
	config.PreferSimpleProtocol = true
	sqldb := stdlib.OpenDB(*config)
	Db := bun.NewDB(sqldb, pgdialect.New())
	return Db
}
