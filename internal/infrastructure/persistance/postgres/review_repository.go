package postgres

import (
	"database/sql"

	"github.com/vandannandwana/MovieReviewApp/internal/domain"
)

type postgresReviewRepository struct {
	db *sql.DB
}

func NewPostgreReviewRepository (db *sql.DB) (domain.ReviewRepository){
	return &postgresReviewRepository{db : db}
}

func (r *postgresReviewRepository) New(review *domain.Review) error{

	return nil
}
func (r *postgresReviewRepository) GetReviewById(id int64) (*domain.Review, error){

	return nil, nil
}
func (r *postgresReviewRepository) GetReviewByMovieId(id int64) (*domain.Review, error){

	return nil, nil
}

func (r *postgresReviewRepository) GetReviewByUserEmailId(email string) (*domain.Review, error){

	return nil, nil
}
func (r *postgresReviewRepository) Update(review *domain.Review) error{

	return nil
}
func (r *postgresReviewRepository) Delete(reviewId int64) error{

	return nil
}