package usecase

import (
	"time"

	"github.com/vandannandwana/MovieReviewApp/internal/delivery/http/dto"
	"github.com/vandannandwana/MovieReviewApp/internal/domain"
)

type ReviewService interface {
	CreateReview(review *domain.Review) error
	GetReviewById(id int64) (*dto.ReviewResponse, error)
	GetReviewByMovieId(id int64) ([]dto.ReviewResponse, error)
	GetReviewByUserEmailId(email string) ([]dto.ReviewResponse, error)
	UpdateReview(review *domain.Review, reviewId int64) error
	DeleteReview(reviewId int64) error
}

type reviewService struct {
	reviewRepo domain.ReviewRepository
}

func NewReviewService(reviewRepo domain.ReviewRepository) ReviewService {
	return &reviewService{reviewRepo: reviewRepo}
}

func (s *reviewService) CreateReview(review *domain.Review) error {

	review.PublishedOn = time.Now()
	review.LastEditOn = time.Now()


	err := s.reviewRepo.New(review)

	if err != nil{
		return err
	}

	return nil
}
func (s *reviewService) GetReviewById(id int64) (*dto.ReviewResponse, error) {

	review, err := s.reviewRepo.GetReviewById(id)

	if err != nil{
		return  nil, err
	}

	return review, nil
}
func (s *reviewService) GetReviewByMovieId(movieId int64) ([]dto.ReviewResponse, error) {

	reviews, err := s.reviewRepo.GetReviewByMovieId(movieId)

	if err != nil{
		return nil, err
	}

	return reviews, nil
}
func (s *reviewService) GetReviewByUserEmailId(emailId string) ([]dto.ReviewResponse, error) {
	reviews, err := s.reviewRepo.GetReviewByUserEmailId(emailId)

	if err != nil{
		return nil, err
	}

	return reviews, nil
}
func (s *reviewService) UpdateReview(review *domain.Review, reviewId int64) error {

	err := s.reviewRepo.Update(review, reviewId)

	if err !=nil{
		return err
	}

	return nil
}
func (s *reviewService) DeleteReview(reviewId int64) error {

	err := s.reviewRepo.Delete(reviewId)

	if err != nil{
		return err
	}

	return nil
}
