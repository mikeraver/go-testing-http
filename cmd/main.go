package main

import (
    "github.com/mikeraver/go-testing-http/api"
    "log"
    "net/http"
    "time"
)

func main() {
	log.Println("registering http handlers")
    customerApi := api.NewCustomerApi()
    customerApi.RegisterHandlers()

    log.Println("starting server on port 8080")

    s := &http.Server {
        ReadHeaderTimeout:20 * time.Second,
        ReadTimeout: 1 * time.Minute,
        WriteTimeout: 2 * time.Minute,
        Addr:":8080",
    }
    log.Fatal(s.ListenAndServe())
}
