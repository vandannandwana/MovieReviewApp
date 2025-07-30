package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/vandannandwana/MovieReviewApp/internal/delivery/http/dto"
	"github.com/vandannandwana/MovieReviewApp/internal/usecase"
	"github.com/vandannandwana/MovieReviewApp/internal/utils/response"
)

type MovieHandler struct {
	movieService usecase.MovieService
}

func NewMovieHandler(movieService usecase.MovieService) *MovieHandler {
	return &MovieHandler{movieService: movieService}
}

func (h *MovieHandler) CreateMovie(c *gin.Context) {

	var movie dto.CreateMovieRequest

	if err := c.ShouldBind(&movie); err != nil {

		c.JSON(http.StatusBadRequest, "Invalid request payload"+err.Error())
		fmt.Println(movie)
		return
	}

	err := h.movieService.CreateMovie(&movie)

	if err != nil {
		c.JSON(http.StatusBadRequest, "faild tp create movie"+err.Error())
	}

	c.JSON(http.StatusCreated, gin.H{"Status": "Movie Created Successfully"})

}

func (h *MovieHandler) GetMovieById(c *gin.Context) {

	var _movieId = c.Param("id")

	if _movieId == "" {
		c.JSON(http.StatusBadRequest, "MovieId Missing")
		return
	}

	movieId, err := strconv.ParseInt(_movieId, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Inputs!")
		return
	}

	movieRes, err := h.movieService.GetMovieById(movieId)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("no element found with the id %s", _movieId)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movieRes)

}
func (h *MovieHandler) UpdateMovie(c *gin.Context) {

	var _movieId = c.Param("id")

	if _movieId == "" {
		c.JSON(http.StatusBadRequest, "MovieId Missing")
		return
	}

	movieId, err := strconv.ParseInt(_movieId, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Inputs!")
		return
	}

	var updateMovieDto dto.UpdateMovieRequest

	err = c.ShouldBindJSON(&updateMovieDto)

	if err != nil {

		if errors.Is(err, io.EOF) {
			c.JSON(http.StatusBadRequest, StandardError("empty body"))
			return
		}

		if err := validator.New().Struct(updateMovieDto); err != nil {
			validationErrs := err.(validator.ValidationErrors)
			c.JSON(http.StatusBadRequest, response.ValidationError(validationErrs))
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.movieService.UpdateMovie(&updateMovieDto, movieId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Status": "Movie Updated Successfully"})

}
func (h *MovieHandler) DeleteMovie(c *gin.Context) {

	var _movieId = c.Param("id")
	var userEmail = c.Param("email")

	if userEmail == "" {
		c.JSON(http.StatusBadRequest, "Not a valid user")
	}

	if _movieId == "" {
		c.JSON(http.StatusBadRequest, "MovieId Missing")
		return
	}

	movieId, err := strconv.ParseInt(_movieId, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Inputs!")
		return
	}

	err = h.movieService.DeleteMovie(movieId, userEmail)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Status": "Movie Deleted Successfully"})

}
