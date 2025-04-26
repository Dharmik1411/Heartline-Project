package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    ConnectDatabase()

    r := mux.NewRouter()

    r.HandleFunc("/register", Register).Methods("POST")
    r.HandleFunc("/login", Login).Methods("POST")
    
    secured := r.PathPrefix("/").Subrouter()
    secured.Use(AuthMiddleware)
    secured.HandleFunc("/profile", GetProfile).Methods("GET")
    secured.HandleFunc("/profile", UpdateProfile).Methods("PATCH")

    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}

