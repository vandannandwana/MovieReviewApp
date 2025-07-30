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
	"github.com/vandannandwana/MovieReviewApp/internal/domain"
	"github.com/vandannandwana/MovieReviewApp/internal/usecase"
	"github.com/vandannandwana/MovieReviewApp/internal/utils/response"
)

type ReviewHandler struct {
	reviewService usecase.ReviewService
}

func NewReviewHandler(reviewService usecase.ReviewService) *ReviewHandler {
	return &ReviewHandler{reviewService: reviewService}
}

func (h *ReviewHandler) CreateReview(c *gin.Context) {

	var review domain.Review

	if err := c.ShouldBindBodyWithJSON(&review); err != nil {

		if errors.Is(err, io.EOF) {
			c.JSON(http.StatusBadRequest, StandardError("empty body"))
			return
		}

		if err := validator.New().Struct(review); err != nil {
			validationErrs := err.(validator.ValidationErrors)
			c.JSON(http.StatusBadRequest, response.ValidationError(validationErrs))
			return
		}

		c.JSON(http.StatusBadRequest, StandardError(err.Error()))
		return
	}

	err := h.reviewService.CreateReview(&review)

	if err != nil {

		if err.Error() == "pq: insert or update on table \"reviews\" violates foreign key constraint \"fk_user\"" {
			c.JSON(http.StatusInternalServerError, StandardError("Register before reviewing the movie"))
			return
		}

		if err.Error() == "pq: new row for relation \"reviews\" violates check constraint \"reviews_rating_check\""{
			c.JSON(http.StatusInternalServerError, StandardError("Keep the rating between 1 to 5 (without any decimal value)."))
			return
		}

		c.JSON(http.StatusInternalServerError, StandardError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Review Created Successfully"})

}
func (h *ReviewHandler) GetReviewById(c *gin.Context) {

	var _reviewId = c.Param("id")

	if _reviewId == "" {
		c.JSON(http.StatusBadRequest, "ReviewId Missing")
		return
	}

	reviewId, err := strconv.ParseInt(_reviewId, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Inputs!")
		return
	}

	reviewRes, err := h.reviewService.GetReviewById(reviewId)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("no element found with the id %s", _reviewId)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviewRes)

}
func (h *ReviewHandler) GetReviewByMovieId(c *gin.Context) {

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

	reviews, err := h.reviewService.GetReviewByMovieId(movieId)

	if err != nil {

		c.JSON(http.StatusInternalServerError, StandardError(err.Error()))
	}

	c.JSON(http.StatusOK, reviews)

}
func (h *ReviewHandler) GetReviewByUserEmailId(c *gin.Context) {

	var emailId = c.Param("id")

	if emailId == "" {
		c.JSON(http.StatusBadRequest, "EmailId Missing")
		return
	}

	reviews, err := h.reviewService.GetReviewByUserEmailId(emailId)

	if err != nil {

		c.JSON(http.StatusInternalServerError, StandardError(err.Error()))
	}

	c.JSON(http.StatusOK, reviews)

}
func (h *ReviewHandler) UpdateReview(c *gin.Context) {

	var _reviewId = c.Param("id")

	if _reviewId == "" {
		c.JSON(http.StatusBadRequest, "ReviewId Missing")
		return
	}

	reviewId, err := strconv.ParseInt(_reviewId, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Inputs!")
		return
	}

	var updateReview domain.Review

	err = c.ShouldBindJSON(&updateReview)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.reviewService.UpdateReview(&updateReview, reviewId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, StandardError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Review Update Successfully"})

}
func (h *ReviewHandler) DeleteReview(c *gin.Context) {

	var _reviewId = c.Param("id")
	var userEmail = c.Param("email")

	if userEmail == "" {
		c.JSON(http.StatusBadRequest, "Not a valid user")
	}

	if _reviewId == "" {
		c.JSON(http.StatusBadRequest, "ReviewId Missing")
		return
	}

	reviewId, err := strconv.ParseInt(_reviewId, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Inputs!")
		return
	}

	err = h.reviewService.DeleteReview(reviewId, userEmail)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "review deleted successfully"})

}

func StandardError(errMsg string) any {

	return gin.H{"error": errMsg}

}
