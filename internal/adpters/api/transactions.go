package api

import (
	"fristTry/internal/Repository"
	"net/http"
)

func GetFunction(w http.ResponseWriter, r *http.Request) {
	transactions := Repository.Transactionimplem{}
	transactions.GetAllTransactions(w, r)

}

func Createtransaction(w http.ResponseWriter, r *http.Request) {
	transaction := Repository.Transactionimplem{}
	transaction.CreateNewTransaction(w, r)
}
