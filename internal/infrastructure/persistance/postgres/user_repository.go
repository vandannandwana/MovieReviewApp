package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/vandannandwana/MovieReviewApp/internal/domain"
)

type postgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository (db *sql.DB) (domain.UserRepository){
	return &postgresUserRepository{db: db}
}

func (r *postgresUserRepository) New(user *domain.User) error{
	query := "INSERT INTO users (name, email, password, bio, gender, joined_on, profile_picture) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := r.db.Exec(query, user.Name, user.Email, user.Password, user.Bio, user.Gender, user.JoinedOn, user.ProfilePicture)

	if err != nil{
		return err
	}
	return nil
}

func (r *postgresUserRepository) GetByEmail(email string) (*domain.User, error){

	stmt, err := r.db.Prepare("SELECT name, email, password, bio, joined_on, profile_picture, gender FROM users WHERE email = $1 LIMIT 1")

	if err != nil{
		return nil, err
	}

	defer stmt.Close()

	var user domain.User

	err = stmt.QueryRow(email).Scan(&user.Name, &user.Email, &user.Password, &user.Bio, &user.JoinedOn, &user.ProfilePicture, &user.Gender)

	if err != nil{
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, fmt.Errorf("query Error: %w", err)
	}

	return &user, nil
}
func (r *postgresUserRepository) Update(user *domain.User, email string) error{

	stmt, err := r.db.Prepare("UPDATE table SET name = $1, email = $2, password = $3, bio = $4, gender = $5, profile_picture = $6 WHERE email = $7")


	if err != nil{
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Password, user.Bio, user.Gender, user.ProfilePicture, email)

	if err != nil{
		if err == sql.ErrNoRows {
			return fmt.Errorf("no user found with the email id %s", fmt.Sprint(email))
		}
		return fmt.Errorf("query Error: %w", err)
	}

	return nil

}
func (r *postgresUserRepository) Delete(email string) error{

	stmt, err := r.db.Prepare("SELECT * FROM users WHERE email = $1")

	if err != nil{
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(email)

	if err != nil{
		if err == sql.ErrNoRows{
			return fmt.Errorf("no user found with the email id: %v", err)
		}else{
			return err
		}
	}

	stmt, err = r.db.Prepare("DELETE FROM users WHERE email = $1")

	if err != nil{
		return err
	}

	_, err = stmt.Exec(email)

	if err != nil{
		return err
	}

	return nil
}