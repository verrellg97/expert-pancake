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
	GetUserSecurityQuestion(w http.ResponseWriter, r *http.Request) error
	PostUserSecurityAnswer(w http.ResponseWriter, r *http.Request) error
	UpsertUserAddress(w http.ResponseWriter, r *http.Request) error

	UpdateUser(w http.ResponseWriter, r *http.Request) error
	UpdateUserPassword(w http.ResponseWriter, r *http.Request) error
}

type User struct {
	AccountId   string `json:"account_id" validate:"required"`
	FullName    string `json:"full_name" validate:"required"`
	Nickname    string `json:"nickname" validate:"required"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type Location struct {
	Province    string `json:"province" validate:"required"`
	Regency     string `json:"regency" validate:"required"`
	District    string `json:"district" validate:"required"`
	FullAddress string `json:"full_address" validate:"required"`
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
	User
}

type LoginRequest struct {
	AccountId string `json:"account_id" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken           string            `json:"access_token"`
	AccessTokenExpiresAt  time.Time         `json:"access_token_expires_at"`
	RefreshToken          string            `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time         `json:"refresh_token_expires_at"`
	User                  LoginUserResponse `json:"user"`
}

type LoginUserResponse struct {
	User
	Location Location `json:"location"`
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

type UpdateUserPasswordRequest struct {
	AccountId        string `json:"account_id" validate:"required"`
	OldPassword      string `json:"old_password"`
	NewPassword      string `json:"new_password" validate:"required"`
	IsForgotPassword bool   `json:"is_forgot_password"`
}

type UpdateUserPasswordResponse struct {
	Message string `json:"message"`
}

type GetUserSecurityQuestionRequest struct {
	AccountId string `json:"account_id" validate:"required"`
}

type GetUserSecurityQuestionResponse struct {
	Question string `json:"question"`
}

type PostUserSecurityAnswerRequest struct {
	AccountId string `json:"account_id" validate:"required"`
	Question  string `json:"question" validate:"required"`
	Answer    string `json:"answer" validate:"required"`
}

type PostUserSecurityAnswerResponse struct {
	Message string `json:"message"`
}
