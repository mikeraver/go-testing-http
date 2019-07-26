package api

import (
	"encoding/json"
	"fmt"
	"github.com/mikeraver/go-testing-http/model"
	"github.com/mikeraver/go-testing-http/repository"
	"github.com/mikeraver/go-testing-http/service"
	"log"
	"net/http"
)

type CustomerApi struct {
	service *service.CustomerService
}

func (c *CustomerApi) Init() {
	fmt.Println("Initializing Customer API")
	repo := repository.NewRepository()
	c.service = service.NewService(repo)
}

func(c *CustomerApi) CreateCustomerHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	customer := model.Customer{}
	err := decoder.Decode(&customer)
	if err != nil {
		log.Fatalf("failed to marshal the customer: %v", err)
	}

	newCustomer := c.service.CreateCustomer(customer)
	//CustomerApi{}.Service
	//cid, err := uuid.NewRandom()
	//if err != nil {
	//	log.Fatalf("failed to generate a uuid: %v", err)
	//}

	//customer.Id = cid.String()

	encoder := json.NewEncoder(w)
	err = encoder.Encode(newCustomer)
	if err != nil {
		log.Fatalf("failed to encode the response: %v", err)
	}
}
