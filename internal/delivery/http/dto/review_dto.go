package dto

import "time"

type CreateReviewRequest struct {
	MovieId     int64  `json:"movie_id"`
	UserId      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
	IsSpoiler   bool   `json:"is_spoiler"`
}

type UpdateReviewRequest struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Rating      int    `json:"rating,omitempty"`
	IsSpoiler   bool   `json:"is_spoiler,omitempty"`
}

type ReviewResponse struct {
	ReviewId    int64     `json:"review_id"`
	MovieId     int64     `json:"movie_id"`
	UserId      string    `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Rating      int       `json:"rating"`
	Likes       int64     `json:"likes"`
	DisLikes    int64     `json:"dislikes"`
	PublishOn   time.Time `json:"publish_on"`
	LastEditOn  time.Time `json:"last_edit_on"`
	IsSpoiler   bool      `json:"is_spoiler"`
}
