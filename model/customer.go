package model

import "fmt"

type Customer struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (c Customer) String() string {
	return fmt.Sprintf("Customer [Id=%v, FirstName=%v, LastName=%v, Email=%v]",
		c.Id, c.FirstName, c.LastName, c.Email)
}
