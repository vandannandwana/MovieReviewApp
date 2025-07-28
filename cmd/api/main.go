package main

import (
	// "fmt"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/vandannandwana/MovieReviewApp/internal/delivery/http"
	"github.com/vandannandwana/MovieReviewApp/internal/delivery/http/handler"
	"github.com/vandannandwana/MovieReviewApp/internal/infrastructure/persistance"
	"github.com/vandannandwana/MovieReviewApp/internal/infrastructure/persistance/postgres"
	"github.com/vandannandwana/MovieReviewApp/internal/usecase"
)

func main() {

	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "postgres"
	dbPassword := "vandan"
	dbName := "postgres"
	httpPort := ":8082"

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" || httpPort == "" {
		log.Fatal("Missing environment variables for database or HTTP port. Ensure DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, HTTP_PORT are set.")
	}

	db, err := sql.Open("postgres", "postgres://postgres:vandan@localhost/postgres?sslmode=disable")

	if err != nil {
		log.Fatalf("Error Connectong to the database: %v", err)
	}


	userDatabase, err := persistance.New(db)
	if err !=nil{
		log.Fatal(err.Error())
	}


	defer func(){
		if err := userDatabase.Db.Close(); err != nil{
			log.Printf("Error closing the Database: %s", err.Error())
		}
	}()

	if err := userDatabase.Db.Ping(); err != nil{
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Database connected successfully")

	//Initialize Repos

	userRepo := postgres.NewPostgresUserRepository(userDatabase.Db)

	//Initialize Services

	userService := usecase.NewUserService(userRepo)

	//Initialize Handlers

	userHandler := handler.NewUserHandler(userService)

	router := http.NewRouter(*userHandler)

	log.Printf("Starting HTTP server on port %s...", httpPort)
	if err := router.Run(httpPort); err != nil{
		log.Fatalf("Failed to start server: %v", err)
	}

}