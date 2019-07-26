package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mikeraver/go-testing-http/model"
	"log"
)

type CustomerRepository struct {}

func NewRepository() *CustomerRepository {
	return &CustomerRepository{}
}

func (c CustomerRepository) Insert(customer model.Customer) string {
	fmt.Printf("Inserting new customer record: name=[%v, %v]", customer.LastName, customer.FirstName)
	id, err := uuid.NewRandom()
	if err != nil {
		log.Fatal("All is lost")
	}
	return id.String()
}
