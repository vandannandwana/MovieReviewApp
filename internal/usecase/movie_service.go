package usecase

import (
	"fmt"

	"github.com/vandannandwana/MovieReviewApp/internal/delivery/http/dto"
	"github.com/vandannandwana/MovieReviewApp/internal/domain"
)

type MovieService interface {
	CreateMovie(movie *dto.CreateMovieRequest) error
	GetMovieById(id int64) (*dto.MovieResponse, error)
	UpdateMovie(movie *dto.UpdateMovieRequest, movieId int64) error
	DeleteMovie(movieId int64) error
}

type movieService struct{
	movieRepo domain.MovieRepository
}

func NewMovieService(movieRepo domain.MovieRepository) MovieService{
	return &movieService{movieRepo: movieRepo}
}

func (s *movieService) CreateMovie(movieDto *dto.CreateMovieRequest) error{

	var movie = domain.Movie{
		UserEmail: movieDto.UserEmail, 
		Title: movieDto.Title, 
		Description: movieDto.Description, 
		ReleasedOn: movieDto.ReleaseOn,
		Images: movieDto.Images,
		Videos: movieDto.Videos,
		Genres: movieDto.Genres,
		Directors: movieDto.Directors,
		Writers: movieDto.Writes,
		Casts: movieDto.Casts,
		AverageRatings: 0,
		OriginCountry: movieDto.OriginCountry,
		Languages: movieDto.Languages,
		ProductionCompanies: movieDto.ProductionCompanies,
		Budget: movieDto.Budget,
		Runtime: movieDto.Runtime,
	}

	err := s.movieRepo.New(&movie)

	if err != nil{
		return err
	}

	return nil
}
func (s *movieService) GetMovieById(id int64) (*dto.MovieResponse, error){

	movieRes, err := s.movieRepo.GetMovieById(id)

	if err != nil{
		return nil, err
	}

	return movieRes, nil
}
func (s *movieService) UpdateMovie(movieDto *dto.UpdateMovieRequest, movieId int64) error{

	fmt.Println(movieDto)


	var movie = domain.Movie{
		Title: movieDto.Title, 
		Description: movieDto.Description, 
		ReleasedOn: movieDto.ReleaseOn,
		Images: movieDto.Images,
		Videos: movieDto.Videos,
		Genres: movieDto.Genres,
		Directors: movieDto.Directors,
		Writers: movieDto.Writes,
		Casts: movieDto.Casts,
		AverageRatings: 0,
		Languages: movieDto.Languages,
		ProductionCompanies: movieDto.ProductionCompanies,
		Budget: movieDto.Budget,
		Runtime: movieDto.Runtime,
	}


	err := s.movieRepo.Update(&movie, movieId)

	if err != nil{
		return err
	}

	return nil
}
func (s *movieService) DeleteMovie(movieId int64) error{

	err := s.movieRepo.Delete(movieId)

	if err != nil{
		return err
	}

	return nil
}

