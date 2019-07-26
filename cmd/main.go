package main

import (
    "github.com/mikeraver/go-testing-http/api"
    "log"
    "net/http"
)

func main() {
	log.Println("registering http handlers")
    customerApi := api.NewCustomerApi()
    customerApi.RegisterHandlers()
    //http.HandleFunc("/customer", customerApi.CreateCustomerHandler)

    log.Println("starting server on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
