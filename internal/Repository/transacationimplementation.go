package Repository

import (
	"context"
	"encoding/json"
	"fristTry/config"
	"fristTry/internal/Models"
	"fristTry/internal/adpters/db"
	kafka_installation "fristTry/internal/adpters/kafka-installation"
	"github.com/uptrace/bun"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type Transactionimplem struct {
}

func Open_connection() *bun.DB {
	config := config.LoadConfig()
	Db := db.CreateDatabase(config)
	return Db
}

func (t Transactionimplem) GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := Open_connection()
	Ctx := context.Background()
	var Transaction []Models.Tranactions

	err := db.NewSelect().Model(&Transaction).Scan(Ctx)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(Transaction)
}
func (t Transactionimplem) CreateNewTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var trans Models.Tranactions
	db := Open_connection()
	json.NewDecoder(r.Body).Decode(&trans)
	transaction := &trans

	_, err := db.NewInsert().Model(transaction).Exec(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	err1 := validate.Struct(&trans)
	if err1 != nil {
		log.Fatal(err)
		panic("there is something wrong")

	}
	kafka_installation.Produce(&trans)
	json.NewEncoder(w).Encode(trans)
}
