package persistance

import "database/sql"

type Postgre struct{
	Db *sql.DB
}

func New(db *sql.DB) (*Postgre, error) {

	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users(
	user_id SERIAL PRIMARY KEY,
	name TEXT,
	email TEXT,
	password TEXT,
	bio TEXT,
	gender TEXT,
	joined_on DATE,
	profile_picture TEXT
	)`)

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS movies(
	movie_id SERIAL PRIMARY KEY,
	title TEXT,
	user_email TEXT,
	description TEXT,
	released_on DATE,
	images TEXT[],
	videos TEXT[],
	genres TEXT[],
	directors TEXT[],
	writers TEXT[],
	casts TEXT[],
	avg_rating INTEGER,
	origin_country TEXT,
	languages TEXT[],
	production_companies TEXT[],
	budget INTEGER,
	runtime TEXT
	)`)

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS reviews(
	review_id SERIAL PRIMARY KEY,
	movie_id INTEGER,
	user_email TEXT,
	title TEXT,
	description TEXT,
	rating INTEGER,
	likes INTEGER,
	dislikes INTEGER,
	published_on TIMESTAMPTZ,
	last_edit_on TIMESTAMPTZ,
	is_spoiler BOOLEAN
	)`)

	if err != nil {
		return nil, err
	}

	return &Postgre{
		Db: db,
	}, nil

}