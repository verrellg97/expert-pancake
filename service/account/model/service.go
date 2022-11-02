package model

import (
	"net/http"
	"time"
)

type AccountService interface {
	HelloWorld(w http.ResponseWriter, r *http.Request) error
	HelloError(w http.ResponseWriter, r *http.Request) error

	Register(w http.ResponseWriter, r *http.Request) error

	CheckPhoneNumber(w http.ResponseWriter, r *http.Request) error
	PostOtp(w http.ResponseWriter, r *http.Request) error
	Login(w http.ResponseWriter, r *http.Request) error

	GetDefaultSecurityQuestions(w http.ResponseWriter, r *http.Request) error
	UpsertUserAddress(w http.ResponseWriter, r *http.Request) error

	UpdateUser(w http.ResponseWriter, r *http.Request) error
	UpdateUserPassword(w http.ResponseWriter, r *http.Request) error
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
	Email            string `json:"email"`
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

type PhoneNumberRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type PhoneNumberResponse struct {
	AccountId string `json:"account_id"`
}

type PostOtpRequest struct {
	Otp string `json:"otp" validate:"required"`
}

type PostOtpResponse struct {
	Message string `json:"message"`
}

type GetDefaultSecurityQuestionsResponse struct {
	Questions []string `json:"questions"`
}

type UpsertUserAddressRequest struct {
	AccountId string `json:"account_id" validate:"required"`
	Location
}

type UpsertUserAddressResponse struct {
	Message string `json:"message"`
}

type UpdateUserRequest struct {
	AccountId   string   `json:"account_id" validate:"required"`
	FullName    string   `json:"full_name" validate:"required"`
	Nickname    string   `json:"nickname" validate:"required"`
	Email       string   `json:"email"`
	PhoneNumber string   `json:"phone_number" validate:"required"`
	Location    Location `json:"location"`
}

type UpdateUserResponse struct {
	AccountId   string   `json:"account_id"`
	FullName    string   `json:"full_name"`
	Nickname    string   `json:"nickname"`
	Email       string   `json:"email"`
	PhoneNumber string   `json:"phone_number"`
	Location    Location `json:"location"`
}

type Location struct {
	Province    string `json:"province"`
	Regency     string `json:"regency"`
	District    string `json:"district"`
	FullAddress string `json:"full_address"`
}

type UpdateUserPasswordRequest struct {
	AccountId   string `json:"account_id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type UpdateUserPasswordResponse struct {
	Message string `json:"message"`
}
