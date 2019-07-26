package service

import (
	"fmt"
	"github.com/mikeraver/go-testing-http/model"
	"github.com/mikeraver/go-testing-http/repository"
)

type CustomerService struct {
	customerRepository *repository.CustomerRepository
}

func NewService(repository *repository.CustomerRepository) *CustomerService {
	service := CustomerService{repository }
	return &service
}

func (c CustomerService) CreateCustomer(customer model.Customer) model.Customer {
	fmt.Println("Creating customer from the service")
	id := c.customerRepository.Insert(customer)
	customer.Id = id
	return customer
}
