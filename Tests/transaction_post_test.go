package Tests

import (
	"bytes"
	"fristTry/internal/adpters/api"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateEntry(t *testing.T) {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	var jsonStr = []byte(`{
        "Id": "d9a74f8b-c9cd-4493-9f7f-89fd871c9506",
        "Amount": 90000,
        "Currency": "Usd",
        "Createdat": "1/6/2022",
        "Status" : "false"
    }`)
	req, err := http.NewRequest("POST", "/api/trans", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.Post("/api/trans", api.Createtransaction)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	var requst = rr.Body.String()
	requst = strings.TrimSuffix(requst, "\n")
	expected := `{"Id":"d9a74f8b-c9cd-4493-9f7f-89fd871c9506","Amount":90000,"Currency":"Usd","Createdat":"1/6/2022","Status":"false"}`
	x := strings.Compare(expected, requst)
	if x == 1 || x == -1 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
