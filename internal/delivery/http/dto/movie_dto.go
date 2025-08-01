package dto

import "time"

type CreateMovieRequest struct {
	Title               string    `json:"title" binding:"required"`
	UserEmail           string    `json:"email" binding:"required"`
	Description         string    `json:"description"`
	ReleaseOn           time.Time `json:"release_on"`
	Images              []string  `json:"images"`
	Videos              []string  `json:"videos"`
	Genres              []string  `json:"genres"`
	Directors           []string  `json:"directors"`
	Writes              []string  `json:"writes"`
	Casts               []string  `json:"casts"`
	OriginCountry       string    `json:"origin_country"`
	Languages           []string  `json:"languages"`
	ProductionCompanies []string  `json:"production_companies"`
	Budget              float64   `json:"budget"`
	Runtime             string    `json:"runtime"`
}

type UpdateMovieRequest struct {
	Title               string    `json:"title,omitempty"`
	UserEmail           string    `json:"email" binding:"required"`
	Description         string    `json:"description,omitempty"`
	ReleaseOn           time.Time `json:"release_on"`
	Images              []string  `json:"images,omitempty"`
	Videos              []string  `json:"videos,omitempty"`
	Genres              []string  `json:"genres,omitempty"`
	Directors           []string  `json:"directors,omitempty"`
	Writes              []string  `json:"writes,omitempty"`
	Casts               []string  `json:"casts,omitempty"`
	OriginCountry       string    `json:"origin_country,omitempty"`
	Languages           []string  `json:"languages,omitempty"`
	ProductionCompanies []string  `json:"production_companies,omitempty"`
	Budget              float64   `json:"budget,omitempty"`
	Runtime             string    `json:"runtime,omitempty"`
}

type MovieResponse struct {
	MovieId             int64     `json:"movie_id"`
	UserEmail           string    `json:"email"`
	Title               string    `json:"title"`
	Description         string    `json:"description"`
	ReleaseOn           time.Time `json:"release_on"`
	Images              []string  `json:"images"`
	Videos              []string  `json:"videos"`
	Genres              []string  `json:"genres"`
	Directors           []string  `json:"directors"`
	Writers              []string  `json:"writes"`
	Casts               []string  `json:"casts"`
	AverageRatings      float64       `json:"average_ratings"`
	OriginCountry       string    `json:"origin_country"`
	Languages           []string  `json:"languages"`
	ProductionCompanies []string  `json:"production_companies"`
	Budget              float64   `json:"budget"`
	Runtime             string    `json:"runtime"`
	// Reviews will typically be a separate endpoint/DTO for a list
}
