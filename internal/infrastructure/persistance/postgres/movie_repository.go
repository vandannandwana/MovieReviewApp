package postgres

import (
	"database/sql"
	"fmt"

	"github.com/vandannandwana/MovieReviewApp/internal/delivery/http/dto"
	"github.com/vandannandwana/MovieReviewApp/internal/domain"
)

type postgreMovieRepository struct {
	db *sql.DB
}

func NewPostgreMovieRepository (db *sql.DB) (domain.MovieRepository){
	return &postgreMovieRepository{db: db}
}


func (r *postgreMovieRepository) New(movie *domain.Movie) error{

	stmt, err := r.db.Prepare("INSERT INTO movies (user_email, title, description, released_on, images, videos, genres, directors, writers, casts, avg_rating, origin_country, languages, production_companies, budget, runtime) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)")

	if err != nil{
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(movie.UserEmail, movie.Title, movie.Description, movie.ReleasedOn, movie.Images, movie.Videos, movie.Genres, movie.Directors, movie.Writers, movie.Casts, movie.AverageRatings, movie.OriginCountry, movie.Languages, movie.ProductionCompanies, movie.Budget, movie.Runtime)

	if err != nil{
		return err
	}

	return nil
}
func (r *postgreMovieRepository) GetMovieById(id int64) (*dto.MovieResponse, error){

	stmt, err := r.db.Prepare("SELECT movie_id, title, description, released_on, images, videos, genres, directors, writers, casts, avg_rating, origin_country, languages, production_companies, budget, runtime FROM movies WHERE movie_id = $1")

	if err != nil{
		return nil, err
	}

	defer stmt.Close()

	var movie dto.MovieResponse

	err = stmt.QueryRow(id).Scan(&movie.MovieId, &movie.Title, &movie.Description, &movie.ReleaseOn, &movie.Images, &movie.Videos, &movie.Genres, &movie.Directors, &movie.Writers, &movie.Casts, &movie.AverageRatings, &movie.OriginCountry, &movie.Languages, &movie.ProductionCompanies, &movie.Budget, &movie.Runtime)

	if err != nil{
		return nil, err
	}

	return &movie, nil
}
func (r *postgreMovieRepository) Update(movie *domain.Movie, movieId int64) error{

	stmt, err := r.db.Prepare("UPDATE movies SET title = $1, description = $2, images = $3, videos = $4, genres = $5, directors = $6, writers = $7, casts = $8, avg_rating = $9, origin_country = $10, languages = $11, production_companies = $12, budget = $13, runtime = $14 WHERE movie_id = $15")

	if err != nil{
		return err
	}

	defer stmt.Close()

	_, err = r.db.Exec(movie.Title, movie.Description, movie.Images, movie.Videos, movie.Genres, movie.Directors, movie.Writers, movie.Casts, movie.AverageRatings, movie.OriginCountry, movie.Languages, movie.ProductionCompanies, movie.Budget, movie.Runtime, movieId)

	if err != nil{
		return err
	}

	return nil
}
func (r *postgreMovieRepository) Delete(movieId int64) error{

	_, err := r.db.Prepare("SELECT * FROM movies WHERE movie_id = $1")

	if err != nil{
		if err == sql.ErrNoRows{
			return fmt.Errorf("no movie found with the id: %d", movieId)
		}else{
			return err
		}
	}

	query := "DELETE FROM movies WHERE movie_id = $1"

	_, err = r.db.Exec(query, movieId)

	if err != nil{
		return err
	}

	return nil
}
