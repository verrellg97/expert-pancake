package model

import (
	"net/http"
	"time"
)

type AccountService interface {
	HelloWorld(w http.ResponseWriter, r *http.Request) error
	HelloError(w http.ResponseWriter, r *http.Request) error

	Register(w http.ResponseWriter, r *http.Request) error

	Login(w http.ResponseWriter, r *http.Request) error
}

type UserResponse struct {
	Id          string `json:"id"`
	FullName    string `json:"full_name"`
	Nickname    string `json:"nickname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type RegisterRequest struct {
	FullName         string `json:"full_name" validate:"required"`
	Nickname         string `json:"nickname" validate:"required"`
	Email            string `json:"email" validate:"required"`
	PhoneNumber      string `json:"phone_number" validate:"required"`
	Password         string `json:"password" validate:"required"`
	SecurityQuestion string `json:"security_question" validate:"required"`
	SecurityAnswer   string `json:"security_answer" validate:"required"`
}

type RegisterResponse struct {
	UserResponse
}

type LoginRequest struct {
	AccountId string `json:"account_id" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken           string       `json:"access_token"`
	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
	RefreshToken          string       `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
	User                  UserResponse `json:"user"`
}
