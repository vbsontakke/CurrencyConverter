package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
   
)

func main() {
    router := mux.NewRouter()

    // Define routes
    router.HandleFunc("/currencyconverter", ).Methods("GET")
    router.HandleFunc("/fetchCurrentRate", ).Methods("GET")
    
    log.Fatal(http.ListenAndServe(":8080", router))
}
