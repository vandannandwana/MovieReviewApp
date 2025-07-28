package postgres

import (
	"database/sql"

	"github.com/vandannandwana/MovieReviewApp/internal/domain"
	_ "github.com/lib/pq"
	
)

type postgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository (db *sql.DB) (domain.UserRepository){
	return &postgresUserRepository{db: db}
}

func (r *postgresUserRepository) New(user *domain.User) error{
	query := "INSERT INTO users (name, email, password, bio, gender, joinedOn, profilePicture) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := r.db.Exec(query, user.Name, user.Email, user.Password, user.Bio, user.Gender, user.JoinedOn, user.ProfilePicture)

	if err != nil{
		return err
	}
	return nil
}

func (r *postgresUserRepository) GetByEmail(email string) (*domain.User, error){
	return nil, nil
}
func (r *postgresUserRepository) Update(user *domain.User) error{
	return nil
}
func (r *postgresUserRepository) Delete(email string) error{
	return nil
}