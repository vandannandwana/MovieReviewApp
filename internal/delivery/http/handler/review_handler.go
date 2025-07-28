package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vandannandwana/MovieReviewApp/internal/usecase"
)

type ReviewHandler struct {
	reviewService usecase.ReviewService
}

func NewReviewHandler(reviewService usecase.ReviewService) *ReviewHandler {
	return &ReviewHandler{reviewService: reviewService}
}

func (h *ReviewHandler) CreateReview(c *gin.Context) {
}
func (h *ReviewHandler) GetReviewById(c *gin.Context) {
}
func (h *ReviewHandler) GetReviewByMovieId(c *gin.Context) {
}
func (h *ReviewHandler) GetReviewByUserEmailId(c *gin.Context) {
}
func (h *ReviewHandler) UpdateReview(c *gin.Context) {
}
func (h *ReviewHandler) DeleteReview(c *gin.Context) {
}
