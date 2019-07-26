package service

import (
	"github.com/google/uuid"
	"github.com/mikeraver/go-testing-http/model"
	"log"
	"testing"
)

type mockCustomerRepository struct {}
func (c mockCustomerRepository) Insert(customer model.Customer) *model.Customer {
	id, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}

	return &model.Customer {
		Id: id.String(),
		FirstName: "Test",
		LastName: "User",
		Email: "test@test.com",
	}
}

func TestCreateCustomer(t *testing.T) {
	customerService := customerService{}
	customerService.SetCustomerRepository(mockCustomerRepository{})

	expectedCustomer := model.Customer {
		FirstName: "Test",
		LastName: "User",
		Email: "test@test.com",
	}

	actualCustomer := customerService.CreateCustomer(expectedCustomer)
	if &actualCustomer == nil {
		t.Error("customer returned as null")
	}
}
