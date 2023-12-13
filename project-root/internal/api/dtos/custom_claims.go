package dtos

import "github.com/golang-jwt/jwt/v5"

// CustomClaims represents the custom claims in the JWT token.
type CustomClaims struct {
	Name   string `json:"name"`
	Admin  bool   `json:"admin"`
	Claims jwt.RegisteredClaims
}
