package domain

import (
	"time"

	"github.com/vandannandwana/MovieReviewApp/internal/delivery/http/dto"
)

type Movie struct {
	MovieId int64 
	UserEmail           string    `json:"email" binding:"required"`
	Title               string    `json:"title" binding:"required"`
	Description         string    `json:"description" binding:"required"`
	ReleasedOn          time.Time `json:"released_on" binding:"required" time_format:"2006-01-02"`
	Images              []string  `json:"images"`
	Videos              []string  `json:"videos"`
	Genres              []string  `json:"genres"`
	Directors           []string  `json:"directors" binding:"required"`
	Writers              []string  `json:"writers" binding:"required"`
	Casts               []string  `json:"casts" binding:"required"`
	AverageRatings      int       `json:"rating"`
	OriginCountry       string    `json:"origin_country" binding:"required"`
	Languages           []string  `json:"languages" binding:"required"`
	ProductionCompanies []string  `json:"production_company" binding:"required"`
	Budget              float64   `json:"budget"`
	Runtime             string    `json:"runtime"`
}

type MovieRepository interface {
	New(movie *Movie) error
	GetMovieById(id int64) (*dto.MovieResponse, error)
	Update(movie *Movie, movieId int64) error
	Delete(movieId int64) error
}
