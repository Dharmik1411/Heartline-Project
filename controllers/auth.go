package controllers

import (
    "encoding/json"
    "net/http"
    "simple-auth-api/config"
    "simple-auth-api/models"
    "simple-auth-api/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)

    hashedPassword, _ := utils.HashPassword(user.Password)
    user.Password = hashedPassword

    err := config.DB.QueryRow("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id",
        user.Name, user.Email, user.Password).Scan(&user.ID)
    if err != nil {
        http.Error(w, "Error creating user", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
    var user models.User
    var storedPassword string
    json.NewDecoder(r.Body).Decode(&user)

    row := config.DB.QueryRow("SELECT password FROM users WHERE email=$1", user.Email)
    err := row.Scan(&storedPassword)
    if err != nil {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
    }

    if !utils.CheckPasswordHash(user.Password, storedPassword) {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
    }

    token, _ := utils.GenerateJWT(user.Email)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
    authHeader := r.Header.Get("Authorization")
    token := authHeader[len("Bearer "):]

    email, _ := utils.ParseJWT(token)

    var user models.User
    row := config.DB.QueryRow("SELECT id, name, email FROM users WHERE email=$1", email)
    err := row.Scan(&user.ID, &user.Name, &user.Email)
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(user)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
    authHeader := r.Header.Get("Authorization")
    token := authHeader[len("Bearer "):]

    email, _ := utils.ParseJWT(token)

    var updated models.User
    json.NewDecoder(r.Body).Decode(&updated)

    _, err := config.DB.Exec("UPDATE users SET name=$1, password=$2 WHERE email=$3",
        updated.Name, utils.HashPasswordSimple(updated.Password), email)

    if err != nil {
        http.Error(w, "Error updating profile", http.StatusInternalServerError)
        return
    }

    w.Write([]byte("Profile updated successfully"))
}
