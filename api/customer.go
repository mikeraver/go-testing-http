package api

import (
	"encoding/json"
	"fmt"
	"github.com/mikeraver/go-testing-http/model"
	"github.com/mikeraver/go-testing-http/service"
	"log"
	"net/http"
)

type CustomerApi struct {
	customerService service.CustomerService
}

func NewCustomerApi() *CustomerApi {
	return &CustomerApi{}
}

func (c *CustomerApi) RegisterHandlers() {
	http.HandleFunc("/customer", c.createCustomerHandler)
}

func (c *CustomerApi) SetCustomerService(customerService service.CustomerService) {
	c.customerService = customerService
}

func (c *CustomerApi) getCustomerService() service.CustomerService {
	if c.customerService == nil {
		c.customerService = service.NewCustomerService()
	}

	return c.customerService
}

func(c *CustomerApi) createCustomerHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	customer := model.Customer{}
	err := decoder.Decode(&customer)
	if err != nil {
		log.Fatalf("failed to marshal the customer: %v", err)
	}

	newCustomer := c.getCustomerService().CreateCustomer(customer)
	//CustomerApi{}.Service
	//cid, err := uuid.NewRandom()
	//if err != nil {
	//	log.Fatalf("failed to generate a uuid: %v", err)
	//}

	//customer.Id = cid.String()

	fmt.Printf("\nOUTPUT: %v END_OUTPUT\n", newCustomer)

	encoder := json.NewEncoder(w)
	err = encoder.Encode(newCustomer)
	if err != nil {
		log.Fatalf("failed to encode the response: %v", err)
	}
}
