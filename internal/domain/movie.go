package domain

import "time"

type Movie struct {
	MovieId             int64
	Title               string
	Description         string
	ReleaseOn           time.Time
	Images              []string
	Videos              []string
	Genres              []string
	Directors           []string
	Writes              []string
	Casts               []string
	AverageRatings      int
	Reviews             []Review
	OriginCountry       string
	Languages           []string
	ProductionCompanies []string
	Budget              float64
	Runtime             string
}

type MovieRepository interface{
	New(movie *Movie) error
	GetByMovieId(id int64) (*Movie, error)
	Update(movie *Movie) error
	Delete(movie *Movie) error
}
