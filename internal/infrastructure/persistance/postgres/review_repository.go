package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/vandannandwana/MovieReviewApp/internal/delivery/http/dto"
	"github.com/vandannandwana/MovieReviewApp/internal/domain"
)

type postgresReviewRepository struct {
	db *sql.DB
}

func NewPostgreReviewRepository (db *sql.DB) (domain.ReviewRepository){
	return &postgresReviewRepository{db : db}
}

func (r *postgresReviewRepository) New(review *domain.Review) error{
	//Checking the previous review
	query := "SELECT review_id FROM reviews WHERE user_email = $1 AND movie_id = $2 LIMIT 1"
	stmt, err := r.db.Prepare(query)

	if err != nil{
		fmt.Println(err.Error())
	}
	
	var prevReviewId int64

	err = stmt.QueryRow(review.UserEmail, review.MovieId).Scan(&prevReviewId)

	if prevReviewId != 0{
		return fmt.Errorf("already review exists with the email id")
	}

	if !errors.Is(err, sql.ErrNoRows){
		return err
	}

	//Inserting new review
	query = "INSERT INTO reviews (movie_id, user_email, title, description, rating, likes, dislikes, published_on, last_edit_on, is_spoiler) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

	_, err = r.db.Exec(query, review.MovieId,review.UserEmail, review.Title, review.Description, review.Rating, review.Likes, review.DisLikes, review.PublishedOn, review.LastEditOn, review.IsSpoiler)

	if err != nil{
		return err
	}

	return nil
}
func (r *postgresReviewRepository) GetReviewById(id int64) (*dto.ReviewResponse, error){

	stmt, err := r.db.Prepare("SELECT review_id, movie_id, user_email, title, description, rating, likes, dislikes, published_on, last_edit_on, is_spoiler FROM reviews WHERE review_id = $1 LIMIT 1")

	if err != nil{
		return nil, err
	}

	var review dto.ReviewResponse

	err = stmt.QueryRow(id).Scan(&review.ReviewId, &review.MovieId, &review.UserId, &review.Title, &review.Description, &review.Rating, &review.Likes, &review.DisLikes, &review.PublishOn, &review.LastEditOn, &review.IsSpoiler)

	if err != nil{
		return nil, err
	}

	return &review, nil
}
func (r *postgresReviewRepository) GetReviewByMovieId(movie_id int64) ([]dto.ReviewResponse, error){

	stmt, err := r.db.Prepare("SELECT review_id, movie_id, user_email, title, description, rating, likes, dislikes, published_on, last_edit_on, is_spoiler FROM reviews WHERE movie_id = $1")

	if err != nil{
		return nil, err
	}

	res, err := stmt.Query(movie_id)

	if err != nil{
		return nil, err
	}

	var reviews []dto.ReviewResponse

	for res.Next(){
		var review dto.ReviewResponse
		err := res.Scan(&review.ReviewId, &review.MovieId, &review.UserId, &review.Title, &review.Description, &review.Rating, &review.Likes, &review.DisLikes, &review.PublishOn, &review.LastEditOn, &review.IsSpoiler)
		if err != nil{
			return nil, err
		}
		reviews = append(reviews, review)
	}

	return reviews, nil
}

func (r *postgresReviewRepository) GetReviewByUserEmailId(email string) ([]dto.ReviewResponse, error){

	stmt, err := r.db.Prepare("SELECT review_id, movie_id, user_email, title, description, rating, likes, dislikes, published_on, last_edit_on, is_spoiler FROM reviews WHERE user_email = $1")

	if err != nil{
		return nil, err
	}

	res, err := stmt.Query(email)

	if err != nil{
		return nil, err
	}

	var reviews []dto.ReviewResponse

	for res.Next(){
		var review dto.ReviewResponse
		err := res.Scan(&review.ReviewId, &review.MovieId, &review.UserId, &review.Title, &review.Description, &review.Rating, &review.Likes, &review.DisLikes, &review.PublishOn, &review.LastEditOn, &review.IsSpoiler)
		if err != nil{
			return nil, err
		}
		reviews = append(reviews, review)
	}

	return reviews, nil
}
func (r *postgresReviewRepository) Update(review *domain.Review, reviewId int64) error{

	query := "UPDATE reviews SET title = $1, description = $2, rating = $3, is_spoiler = $4 WHERE review_id = $5 AND user_email = $6 AND movie_id = $7"

	_, err := r.db.Exec(query, review.Title, review.Description, review.Rating, review.IsSpoiler, reviewId, review.UserEmail, review.MovieId)

	if err != nil{
		return err
	}

	return nil
}
func (r *postgresReviewRepository) Delete(reviewId int64, userEmail string) error{

	stmt, err := r.db.Prepare("SELECT user_email FROM reviews WHERE review_id = $1")

	if err != nil{
		slog.Error("Prepare statement error",slog.String("error: ", err.Error()))
	}

	defer stmt.Close()

	var prevUserEmail string

	err = stmt.QueryRow(reviewId).Scan(&prevUserEmail)

	if err != nil{
		if errors.Is(err, sql.ErrNoRows){
			return fmt.Errorf("no review found with the id: %d", reviewId)
		}else{
			return err
		}
	}

	if prevUserEmail != userEmail{
		return fmt.Errorf("you are not allowed to delete the review, only owners are allowed")
	}

	query := "DELETE FROM reviews WHERE review_id = $1"

	_, err = r.db.Exec(query, reviewId)

	if err != nil{
		return err
	}

	return nil
}