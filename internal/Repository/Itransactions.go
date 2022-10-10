package Repository

import (
	"fristTry/internal/Models"
	"net/http"
)

type TransactionDetalis interface {
	GetAllTransactions(w *http.Request, r *http.Request)
	CreateNewTransaction(w *http.Request, r *http.Request)
	UpdateTransaction(tranactions *Models.Tranactions)
}
