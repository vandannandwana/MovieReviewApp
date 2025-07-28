package domain

import "time"

type Movie struct {
	MovieId             int64
	UserEmail           string `json:"email" binding:"required"`
	Title               string `json:"title" binding:"required"`
	Description         string `json:"description" binding:"required"`
	ReleasedOn          time.Time `json:"releasedOn" binding:"required" time_format:"2006-01-02"`
	Images              []string `json:"images"`
	Videos              []string `json:"videos"`
	Genres              []string `json:"genres"`
	Directors           []string `json:"directors" binding:"required"`
	Writes              []string `json:"writers" binding:"required"`
	Casts               []string `json:"casts" binding:"required"`
	AverageRatings      int `json:"rating"`
	OriginCountry       string `json:"originCountry" binding:"required"`
	Languages           []string `json:"languages" binding:"required"`
	ProductionCompanies []string `json:"productionCompany" binding:"required"`
	Budget              float64 `json:"budget"`
	Runtime             string `json:"runtime"`
}

type MovieRepository interface {
	New(movie *Movie) error
	GetMovieById(id int64) (*Movie, error)
	Update(movie *Movie) error
	Delete(movieId int64) error
}
