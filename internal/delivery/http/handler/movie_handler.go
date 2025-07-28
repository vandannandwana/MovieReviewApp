package handler

import (
	"github.com/gin-gonic/gin"
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

func (h *MovieHandler) CreateMovie(c *gin.Context) {
}

func (h *MovieHandler) GetMovieById(c *gin.Context) {
}
func (h *MovieHandler) UpdateMovie(c *gin.Context) {
}
func (h *MovieHandler) DeleteMovie(c *gin.Context) {
}
