package utils

import (
    "github.com/golang-jwt/jwt/v5"
    "time"
)

var jwtKey = []byte("your_secret_key") 

func GenerateJWT(email string) (string, error) {
    claims := jwt.MapClaims{}
    claims["authorized"] = true
    claims["email"] = email
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func ParseJWT(tokenString string) (string, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        email := claims["email"].(string)
        return email, nil
    } else {
        return "", err
    }
}
