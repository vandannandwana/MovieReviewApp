package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/lib/pq"
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

	_, err = stmt.Exec(
		movie.UserEmail,
		movie.Title, 
		movie.Description, 
		movie.ReleasedOn, 
		pq.Array(movie.Images), 
		pq.Array(movie.Videos), 
		pq.Array(movie.Genres), 
		pq.Array(movie.Directors), 
		pq.Array(movie.Writers), 
		pq.Array(movie.Casts), 
		movie.AverageRatings, 
		movie.OriginCountry, 
		pq.Array(movie.Languages), 
		pq.Array(movie.ProductionCompanies), 
		movie.Budget, 
		movie.Runtime,
	)

	if err != nil{
		return err
	}

	return nil
}
func (r *postgreMovieRepository) GetMovieById(id int64) (*dto.MovieResponse, error){

	stmt, err := r.db.Prepare("SELECT movie_id, user_email, title, description, released_on, images, videos, genres, directors, writers, casts, avg_rating, origin_country, languages, production_companies, budget, runtime FROM movies WHERE movie_id = $1")

	if err != nil{
		return nil, err
	}

	defer stmt.Close()

	var movie dto.MovieResponse

	err = stmt.QueryRow(id).Scan(&movie.MovieId, 
		&movie.UserEmail,
		&movie.Title, 
		&movie.Description, 
		&movie.ReleaseOn, 
		pq.Array(&movie.Images), 
		pq.Array(&movie.Videos), 
		pq.Array(&movie.Genres), 
		pq.Array(&movie.Directors), 
		pq.Array(&movie.Writers), 
		pq.Array(&movie.Casts), 
		&movie.AverageRatings, 
		&movie.OriginCountry, 
		pq.Array(&movie.Languages), 
		pq.Array(&movie.ProductionCompanies),
		&movie.Budget, 
		&movie.Runtime,
	)

	if err != nil{
		return nil, err
	}

	return &movie, nil
}
func (r *postgreMovieRepository) Update(movie *domain.Movie, movieId int64) error{

	fmt.Println(movie)

	stmt, err := r.db.Prepare("UPDATE movies SET title = $1, description = $2, images = $3, videos = $4, genres = $5, directors = $6, writers = $7, casts = $8, origin_country = $9, languages = $10, production_companies = $11, budget = $12, runtime = $13 WHERE movie_id = $14")

	if err != nil{
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		movie.Title, 
		movie.Description, 
		pq.Array(movie.Images), 
		pq.Array(movie.Videos), 
		pq.Array(movie.Genres), 
		pq.Array(movie.Directors), 
		pq.Array(movie.Writers), 
		pq.Array(movie.Casts), 
		movie.OriginCountry, 
		pq.Array(movie.Languages), 
		pq.Array(movie.ProductionCompanies), 
		movie.Budget, 
		movie.Runtime, 
		movieId,
	)

	if err != nil{
		return err
	}

	return nil
}
func (r *postgreMovieRepository) Delete(movieId int64, userEmail string) error{

	stmt, err := r.db.Prepare("SELECT user_email FROM movies WHERE movie_id = $1 LIMIT 1")

	if err != nil{
		slog.Error("Prepare statement error",slog.String("error: ", err.Error()))
	}

	defer stmt.Close()

	var prevUserEmail string

	err = stmt.QueryRow(movieId).Scan(&prevUserEmail)

	if err != nil{
		if errors.Is(err, sql.ErrNoRows){
			return fmt.Errorf("no movie found with the id: %d", movieId)
		}else{
			return err
		}
	}

	if prevUserEmail != userEmail{
		return fmt.Errorf("you are not allowed to delete the movie, only owners are allowed")
	}

	query := "DELETE FROM movies WHERE movie_id = $1"

	_, err = r.db.Exec(query, movieId)

	if err != nil{
		return err
	}

	return nil
}
