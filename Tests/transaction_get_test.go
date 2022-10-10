package Tests

import (
	"fristTry/internal/adpters/api"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetEntries(t *testing.T) {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	req, err := http.NewRequest("GET", "/api/trans", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.Get("/api/trans", api.GetFunction)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `[{"Id":"17b1936f-90c5-4d17-9a09-eb4d589e1083","Amount":5000,"Currency":"Mxn","Createdat":"5/5/2022","Status":true},{"Id":"86d41698-0606-4145-abb1-ce6b7c72eb69","Amount":15000,"Currency":"Usd","Createdat":"10/12/2022","Status":true},{"Id":"eaf9071d-d274-4a02-914c-b2bbca743841","Amount":1560,"Currency":"pezo","Createdat":"10/10/2022","Status":false}]`
	x := strings.Compare(expected, rr.Body.String())
	if x == 1 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
