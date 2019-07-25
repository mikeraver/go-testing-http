package main

import (
    "github.com/mikeraver/go-testing-http/api"
    "log"
    "net/http"
)

func main() {

    http.Handle("/customer", api.CustomerAPI{})

    log.Println("starting server on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
