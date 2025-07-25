package domain

import "time"

type Review struct {
	ReviewId    int64
	MovieId     int64 `json:"movieId" binding:"required"`
	UserEmail   string `json:"userEmail" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Rating      int `json:"rating" binding:"required"`
	Likes       int64 `json:"likes"`
	DisLikes    int64 `json:"dislikes"`
	PublishOn   time.Time
	LastEditOn  time.Time
	IsSpoiler   bool `json:"isSpoiler" binding:"required"`
}

type ReviewRepository interface {
	New(review *Review) error
	GetByReviewById(id int64) (*Review, error)
	GetByReviewByMovieId(id int64) (*Review, error)
	GetByReviewByUserEmailId(email string) (*Review, error)
	Update(review *Review) error
	Delete(reviewId int64) error
}
