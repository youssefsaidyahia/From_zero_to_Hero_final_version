package kafka_installation

import (
	"context"
	"fristTry/config"
	"fristTry/internal/Models"
	"fristTry/internal/adpters/db"
	"github.com/uptrace/bun"
	"log"
)

func OpenConnection() *bun.DB {
	config := config.LoadConfig()
	Db := db.CreateDatabase(config)
	return Db
}
func UpdateTransaction(tranaction *Models.Tranactions) {
	db := OpenConnection()
	tranaction.Status = true
	_, err := db.NewUpdate().Model(tranaction).Where("Id = ?", tranaction.Id).Exec(context.Background())
	if err != nil {
		log.Fatalf("there is something wrong with kafka")
	}
}
