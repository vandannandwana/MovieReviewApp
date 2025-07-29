package domain

import (
	"time"

	"github.com/vandannandwana/MovieReviewApp/internal/delivery/http/dto"
)

type Review struct {
	ReviewId int64
	MovieId     int64  `json:"movieId" binding:"required"`
	UserEmail   string `json:"userEmail" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Rating      int    `json:"rating" binding:"required"`
	Likes       int64  
	DisLikes    int64  
	PublishedOn time.Time
	LastEditOn  time.Time
	IsSpoiler   bool `json:"isSpoiler" binding:"required"`
}

type ReviewRepository interface {
	New(review *Review) error
	GetReviewById(id int64) (*dto.ReviewResponse, error)
	GetReviewByMovieId(id int64) ([]dto.ReviewResponse, error)
	GetReviewByUserEmailId(email string) ([]dto.ReviewResponse, error)
	Update(review *Review, reviewId int64) error
	Delete(reviewId int64) error
}
