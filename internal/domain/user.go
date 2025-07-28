package domain

import "time"

type User struct {
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required"`  // Primary Key
	Password       string `json:"password" binding:"required"`
	Bio            string `json:"bio" binding:"required"`
	Gender         string `json:"gender" binding:"required"`
	JoinedOn       time.Time `json:"joinedOn" binding:"required"`
	ProfilePicture string `json:"profilePicture"`
}

type UserRepository interface{
	New(user *User) error
	GetByEmail(email string) (*User, error)
	Update(user *User, email string) error
	Delete(email string) error
}
