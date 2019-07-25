package api

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/mikeraver/go-testing-http/model"
	"log"
	"net/http"
)

type CustomerAPI struct {}

func (c CustomerAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	customer := model.Customer{}
	err := decoder.Decode(&customer)
	if err != nil {
		log.Fatalf("failed to marshal the customer: %v", err)
	}

	cid, err := uuid.NewRandom()
	if err != nil {
		log.Fatalf("failed to generate a uuid: %v", err)
	}

	customer.Id = cid.String()

	encoder := json.NewEncoder(w)
	err = encoder.Encode(customer)
	if err != nil {
		log.Fatalf("failed to encode the response: %v", err)
	}
}
