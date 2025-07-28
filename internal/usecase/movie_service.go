package usecase

import "github.com/vandannandwana/MovieReviewApp/internal/domain"

type MovieService interface {
	CreateMovie(movie *domain.Movie) error
	GetMovieById(id int64) (*domain.Movie, error)
	UpdateMovie(movie *domain.Movie) error
	DeleteMovie(movieId int64) error
}

type movieService struct{
	movieRepo domain.MovieRepository
}

func NewMovieService(movieRepo domain.MovieRepository) MovieService{
	return &movieService{movieRepo: movieRepo}
}

func (s *movieService) CreateMovie(movie *domain.Movie) error{
	return nil
}
func (s *movieService) GetMovieById(id int64) (*domain.Movie, error){
	return nil, nil
}
func (s *movieService) UpdateMovie(movie *domain.Movie) error{
	return nil
}
func (s *movieService) DeleteMovie(movieId int64) error{
	return nil
}

