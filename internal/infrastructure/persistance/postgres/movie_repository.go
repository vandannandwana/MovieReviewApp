package postgres

import (
	"database/sql"

	"github.com/vandannandwana/MovieReviewApp/internal/domain"
)

type postgreMovieRepository struct {
	db *sql.DB
}

func NewPostgreMovieRepository (db *sql.DB) (domain.MovieRepository){
	return &postgreMovieRepository{db: db}
}


func (r *postgreMovieRepository) New(movie *domain.Movie) error{
	return nil
}
func (r *postgreMovieRepository) GetMovieById(id int64) (*domain.Movie, error){
	return nil, nil
}
func (r *postgreMovieRepository) Update(movie *domain.Movie) error{
	return nil
}
func (r *postgreMovieRepository) Delete(movieId int64) error{
	return nil
}
