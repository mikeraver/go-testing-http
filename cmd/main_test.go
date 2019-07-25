package main

import (
	"github.com/mikeraver/go-testing-http/api"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateCustomer(t *testing.T) {

	const payload = `{
						"first_name":"John",
						"last_name":"Doe",
						"email":"johndoe@test.com"
					}`

	req, err := http.NewRequest("GET", "/customer", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.CreateCustomerHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
