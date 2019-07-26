package service

import (
	"fmt"
	"github.com/mikeraver/go-testing-http/model"
	"github.com/mikeraver/go-testing-http/repository"
)

type CustomerService interface {
	CreateCustomer(customer model.Customer) model.Customer
}

type customerService struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerService() CustomerService {
	return &customerService{}
}

func (c *customerService) SetCustomerRepository(repo repository.CustomerRepository) {
	c.customerRepository = repo
}

func (c *customerService) getCustomerRepository() repository.CustomerRepository {
	if c.customerRepository == nil {
		c.customerRepository = repository.NewCustomerRepository()
	}

	return c.customerRepository
}

func (c *customerService) CreateCustomer(customer model.Customer) model.Customer {
	fmt.Println("Creating customer from the service")
	id := c.getCustomerRepository().Insert(customer)
	customer.Id = id.String()
	return customer
}
