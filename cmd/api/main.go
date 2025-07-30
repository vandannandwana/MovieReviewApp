package main

import (
	"context"
	"database/sql"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
	myhttp "github.com/vandannandwana/MovieReviewApp/internal/delivery/http"
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

	postgresDatabase, err := persistance.New(db)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer func() {
		if err := postgresDatabase.Db.Close(); err != nil {
			log.Printf("Error closing the Database: %s", err.Error())
		}
	}()

	if err := postgresDatabase.Db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Database connected successfully")

	//Initialize Repos

	userRepo := postgres.NewPostgresUserRepository(postgresDatabase.Db)
	movieRepo := postgres.NewPostgreMovieRepository(postgresDatabase.Db)
	reviewRepo := postgres.NewPostgreReviewRepository(postgresDatabase.Db)

	//Initialize Services

	userService := usecase.NewUserService(userRepo)
	movieService := usecase.NewMovieService(movieRepo)
	reviewService := usecase.NewReviewService(reviewRepo, movieRepo)

	//Initialize Handlers

	userHandler := handler.NewUserHandler(userService)
	movieHandler := handler.NewMovieHandler(movieService)
	reviewHandler := handler.NewReviewHandler(reviewService)

	router := myhttp.NewRouter(*userHandler, *movieHandler, *reviewHandler)

	server := http.Server{
		Handler: router.Handler(),
	}

	log.Printf("Starting HTTP server on port %s...", httpPort)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	

	go func() {
		if err := router.Run(httpPort); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-done
	
	slog.Info("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown server", slog.String("error: ", err.Error()))
	}

	slog.Info("Server Shutdown Successfully")

}
