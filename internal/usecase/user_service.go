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

func NewUserService (userRepo domain.UserRepository) UserService{
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

	var hashedPassword []byte

	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

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

	err = s.userRepo.New(user)

	if err != nil{
		return nil, err
	}

	return user, nil
}

func (s *userService) LoginUser(email, password string) (bool, error){
	user, err := s.userRepo.GetByEmail(email)

	if err != nil{
		return false, err
	}

	if user == nil{
		return false, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil{
		if err == bcrypt.ErrMismatchedHashAndPassword{
			return false, errors.New("password not matched")
		}else{
			return false, err
		}
	}

	return true, nil

}


