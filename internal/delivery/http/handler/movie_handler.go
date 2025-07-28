package handler

import (
	"github.com/vandannandwana/MovieReviewApp/internal/domain"
	"github.com/vandannandwana/MovieReviewApp/internal/usecase"
)

/*
CreateMovie(movie *domain.Movie) error
GetMovieById(id int64) (*domain.Movie, error)
UpdateMovie(movie *domain.Movie) error
DeleteMovie(movieId int64) error
*/

type MovieHandler struct {
	movieService usecase.MovieService
}

func NewMovieHandler(movieService usecase.MovieService) *MovieHandler {
	return &MovieHandler{movieService: movieService}
}

func (h *MovieHandler) CreateMovie(movie *domain.Movie) error {
	return nil
}

func (h *MovieHandler) GetMovieById(id int64) (*domain.Movie, error) {
	return nil, nil
}
func (h *MovieHandler) UpdateMovie(movie *domain.Movie) error {
	return nil
}
func (h *MovieHandler) DeleteMovie(movieId int64) error {
	return nil
}
