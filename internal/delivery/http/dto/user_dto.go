package dto

import "time"

type RegisterUserRequest struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Bio            string `json:"bio"`
	Gender         string `json:"gender"`
	ProfilePicture string `json:"profile_picture"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Name           string `json:"name,omitempty"`
	Password       string `json:"password,omitempty"`
	Bio            string `json:"bio,omitempty"`
	Gender         string `json:"gender,omitempty"`
	ProfilePicture string `json:"profile_picture,omitempty"`
}

type UserResponse struct {
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Bio            string    `json:"bio"`
	Gender         string    `json:"gender"`
	JoinedOn       time.Time `json:"joined_on"`
	ProfilePicture string    `json:"profile_picture"`
}

type LoginResponse struct {
	Token string `json:"token"`
}