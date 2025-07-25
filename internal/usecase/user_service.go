package usecase

import (
	"errors"
	"time"

	"github.com/vandannandwana/MovieReviewApp/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(name, email, password, bio, gender, profilePicture string) (*domain.User, error)
	LoginUser(email, password string) (bool, error)
}

type userService struct{
	userRepo domain.UserRepository
}

func NewUserService (userRepo domain.UserRepository) *userService{
	return &userService{userRepo: userRepo}
}

func (s *userService) RegisterUser(name, email, password, bio, gender, profilePicture string) (*domain.User, error){

	existingUser, err := s.userRepo.GetByEmail(email)

	if err != nil{
		return nil, err
	}

	if existingUser != nil{
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil{
		return nil, err
	}

	user := &domain.User{
		Name: name,
		Email: email,
		Password: string(hashedPassword),
		Bio: bio,
		Gender: gender,
		JoinedOn: time.Now(),
		ProfilePicture: profilePicture,
	}
	return user, nil
}

func (s *userService) LoginUser(email, password string) (bool, error){
	user, err := s.userRepo.GetByEmail(email)

	if err != nil{
		return false, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil{
		return false, err
	}

	return true, nil

}


