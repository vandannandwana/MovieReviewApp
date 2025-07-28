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
	PublishedOn   time.Time
	LastEditOn  time.Time
	IsSpoiler   bool `json:"isSpoiler" binding:"required"`
}

type ReviewRepository interface {
	New(review *Review) error
	GetReviewById(id int64) (*Review, error)
	GetReviewByMovieId(id int64) (*Review, error)
	GetReviewByUserEmailId(email string) (*Review, error)
	Update(review *Review) error
	Delete(reviewId int64) error
}
