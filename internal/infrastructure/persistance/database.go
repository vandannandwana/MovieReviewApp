package persistance

import (
	"database/sql"
	"fmt"
)

type Postgre struct{
	Db *sql.DB
}

func New(db *sql.DB) (*Postgre, error) {

	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users(
	user_id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	email TEXT UNIQUE NOT NULL,
	password TEXT NOT NULL,
	bio TEXT,
	gender TEXT,
	joined_on DATE NOT NULL,
	profile_picture TEXT
	)`)

	if err != nil {
		return nil, fmt.Errorf("failed to create users table: %w", err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS movies(
	movie_id SERIAL PRIMARY KEY,
	title TEXT NOT NULL,
	user_email TEXT NOT NULL,
	description TEXT NOT NULL,
	released_on DATE NOT NULL,
	images TEXT[],
	videos TEXT[],
	genres TEXT[],
	directors TEXT[],
	writers TEXT[],
	casts TEXT[],
	avg_rating NUMERIC(2,1) NOT NULL DEFAULT 0.0 CHECK (avg_rating >= 0.0 AND avg_rating <= 5.0),
	origin_country TEXT,
	languages TEXT[],
	production_companies TEXT[],
	budget BIGINT,
	runtime TEXT NOT NULL
	)`)

	if err != nil {
		return nil, fmt.Errorf("failed to create movies table: %w", err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS reviews(
	review_id SERIAL PRIMARY KEY,
	movie_id BIGINT NOT NULL,
	user_email TEXT NOT NULL,
	title TEXT NOT NULL,
	description TEXT,
	rating INTEGER NOT NULL CHECK (rating >= 1 AND rating <= 5),
	likes BIGINT NOT NULL DEFAULT 0,
	dislikes BIGINT NOT NULL DEFAULT 0,
	published_on TIMESTAMPTZ NOT NULL,
	last_edit_on TIMESTAMPTZ NOT NULL,
	is_spoiler BOOLEAN NOT NULL DEFAULT FALSE,


	CONSTRAINT fk_movie FOREIGN KEY (movie_id) REFERENCES movies (movie_id) ON DELETE CASCADE,
	
	CONSTRAINT fk_user FOREIGN KEY (user_email) REFERENCES users (email) ON DELETE CASCADE,

	CONSTRAINT unique_user_movie_review UNIQUE (user_email, movie_id)

	)`)

	if err != nil {
		return nil, fmt.Errorf("failed to create reviews table: %w", err)
	}

	return &Postgre{
		Db: db,
	}, nil

}