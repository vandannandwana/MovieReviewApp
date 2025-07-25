package domain

import "time"

type User struct {
	Name           string
	Email          string // Primary Key
	Password       string
	Bio            string
	Gender         string
	JoinedOn       time.Time
	ProfilePicture string
	Movies         []Movie
	Reviews        []Review
	WatchList      []Movie
}

type UserRepository interface{
	New(user *User) error
	GetByEmail(email string) (*User, error)
	Update(user *User) error
	Delete(user *User) error
}
