package main

import jwt "github.com/dgrijalva/jwt-go"

// ErrorResponse is a struct for sending error message with code
type ErrorResponse struct {
	Code    int
	Message string
}

// SuccessResponse is struct
type SuccessResponse struct {
	Code     int
	Message  string
	Response interface{}
}

// Claims is a struct that will be encoded to a jwt
// jwt.StandardClaims is an embedded type to provide expire time
type Claims struct {
	Email string
	jwt.StandardClaims
}

// RegistrationParams is a struct to read the request body
type RegistrationParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginParams is a struct to read the request body
type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SuccessfulLoginResponse is a struct to send the request response
type SuccessfulLoginResponse struct {
	Email     string
	AuthToken string
}

// UserDetails is a struct used for user details
type UserDetails struct {
	Name     string
	Email    string
	Password string
}
