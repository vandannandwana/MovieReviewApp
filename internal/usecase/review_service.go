package usecase

import (
	"database/sql"
	"fmt"
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
	movieRepo  domain.MovieRepository
}

func NewReviewService(reviewRepo domain.ReviewRepository, movieRepo domain.MovieRepository) ReviewService {
	return &reviewService{reviewRepo: reviewRepo, movieRepo: movieRepo}
}

func (s *reviewService) CreateReview(review *domain.Review) error {

	//check owner of the movie

	movie, err := s.movieRepo.GetMovieById(review.MovieId)

	if err != nil{
		if err == sql.ErrNoRows{
			return fmt.Errorf("no movie present with the movie id: %d", review.MovieId)
		}else{
			return err
		}
	}

	if movie.UserEmail == review.UserEmail{
		return fmt.Errorf("owner can't review their own movie")
	}

	review.PublishedOn = time.Now()
	review.LastEditOn = time.Now()

	err = s.reviewRepo.New(review)

	if err != nil {
		return err
	}

	return nil
}
func (s *reviewService) GetReviewById(id int64) (*dto.ReviewResponse, error) {

	review, err := s.reviewRepo.GetReviewById(id)

	if err != nil {
		return nil, err
	}

	return review, nil
}
func (s *reviewService) GetReviewByMovieId(movieId int64) ([]dto.ReviewResponse, error) {

	reviews, err := s.reviewRepo.GetReviewByMovieId(movieId)

	if err != nil {
		return nil, err
	}

	return reviews, nil
}
func (s *reviewService) GetReviewByUserEmailId(emailId string) ([]dto.ReviewResponse, error) {
	reviews, err := s.reviewRepo.GetReviewByUserEmailId(emailId)

	if err != nil {
		return nil, err
	}

	return reviews, nil
}
func (s *reviewService) UpdateReview(review *domain.Review, reviewId int64) error {

	prevReview, err := s.reviewRepo.GetReviewById(reviewId)

	if err != nil {
		return err
	}
	prevTime := prevReview.PublishOn
	fmt.Println(time.Since(prevTime))

	if time.Since(prevTime) > time.Minute*5 {
		return fmt.Errorf("edit time expired for the review with id: %d", reviewId)
	}

	err = s.reviewRepo.Update(review, reviewId)

	if err != nil {
		return err
	}

	return nil
}
func (s *reviewService) DeleteReview(reviewId int64) error {

	err := s.reviewRepo.Delete(reviewId)

	if err != nil {
		return err
	}

	return nil
}
