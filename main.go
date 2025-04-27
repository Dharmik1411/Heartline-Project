package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "simple-auth-api/config"
    "simple-auth-api/controllers"
    "simple-auth-api/middleware"
)

func main() {
    config.Connect()

    r := mux.NewRouter()

    r.HandleFunc("/register", controllers.Register).Methods("POST")
    r.HandleFunc("/login", controllers.Login).Methods("POST")

    protected := r.PathPrefix("/").Subrouter()
    protected.Use(middleware.JWTAuth)
    protected.HandleFunc("/profile", controllers.GetProfile).Methods("GET")
    protected.HandleFunc("/profile", controllers.UpdateProfile).Methods("PATCH")

    http.ListenAndServe(":8080", r)
}
