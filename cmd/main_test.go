package main

import (
	"encoding/json"
	"github.com/mikeraver/go-testing-http/api"
	"github.com/mikeraver/go-testing-http/model"
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

	customerApi := api.CustomerApi{}
	customerApi.Init()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(customerApi.CreateCustomerHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	response := &model.Customer{}
	decoder := json.NewDecoder(rr.Body)
	if err := decoder.Decode(response); err != nil {
		t.Fatal(err)
	}

	if len(response.Id) == 0 {
		t.Errorf("customer id is invalid: got %v want uuid", response.Id)
	}
}
