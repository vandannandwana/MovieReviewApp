package domain

import "time"

type Review struct {
	ReviewId    int64
	Title       string
	Description string
	Rating      int
	Likes       int64
	DisLikes    int64
	PublishOn   time.Time
	LastEditOn  time.Time
	IsSpoiler   bool
}

type ReviewRepository interface{
	New(review *Review) error
	GetByMovieId(id int64) (*Review, error)
	Update(review *Review) error
	Delete(review *Review) error
}
