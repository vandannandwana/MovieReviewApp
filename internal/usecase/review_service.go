package usecase

import "github.com/vandannandwana/MovieReviewApp/internal/domain"

type ReviewService interface {
	CreateReview(review *domain.Review) error
	GetReviewById(id int64) (*domain.Review, error)
	GetReviewByMovieId(id int64) (*domain.Review, error)
	GetReviewByUserEmailId(email string) (*domain.Review, error)
	UpdateReview(review *domain.Review) error
	DeleteReview(reviewId int64) error
}

type reviewService struct {
	reviewRepo domain.ReviewRepository
}

func NewReviewService(reviewRepo domain.ReviewRepository) ReviewService {
	return &reviewService{reviewRepo: reviewRepo}
}

func (s *reviewService) CreateReview(review *domain.Review) error {
	return nil
}
func (s *reviewService) GetReviewById(id int64) (*domain.Review, error) {
	return nil, nil
}
func (s *reviewService) GetReviewByMovieId(id int64) (*domain.Review, error) {
	return nil, nil
}
func (s *reviewService) GetReviewByUserEmailId(email string) (*domain.Review, error) {
	return nil, nil
}
func (s *reviewService) UpdateReview(review *domain.Review) error {
	return nil
}
func (s *reviewService) DeleteReview(reviewId int64) error {
	return nil
}
