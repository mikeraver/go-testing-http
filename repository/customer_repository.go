package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mikeraver/go-testing-http/model"
	"log"
)

type CustomerRepository interface {
	Insert(customer model.Customer) *model.Customer
}

type customerRepository struct {}

func NewCustomerRepository() *customerRepository {
	return &customerRepository{}
}

func (c customerRepository) Insert(customer model.Customer) *model.Customer {
	fmt.Printf("Inserting new customer record: name=[%v, %v]", customer.LastName, customer.FirstName)
	id, err := uuid.NewRandom()
	if err != nil {
		log.Fatal("All is lost")
	}
	customer.Id = id.String()
	return &customer
}
