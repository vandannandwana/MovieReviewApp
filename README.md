Movie Review Backend API
This project is a robust backend API for a movie review platform, built with Go and designed following Clean Architecture principles. It provides core functionalities for user management, movie information, and user-submitted reviews.
# Features
User Management: Register new users, log in, retrieve, update, and delete user profiles.
Movie Management: Add new movies, fetch movie details by ID, update existing movie information, and delete movies.
Review Management: Submit new reviews for movies, retrieve individual reviews, get all reviews for a specific movie, update review content, and delete reviews.
Review Interactions: Like and dislike reviews.
Data Integrity: Enforces that each user can submit only one review per movie at the database level.
PostgreSQL Database: Uses PostgreSQL for persistent data storage.
Clean Architecture: Structured for high maintainability, testability, and scalability.
# Architecture Overview
This application adheres to Clean Architecture principles, which organize code into distinct layers with strict dependency rules. This approach makes the codebase easier to understand, test, and evolve.
The main layers are:
Domain: Contains the core business entities (User, Movie, Review) and their interfaces (e.g., UserRepository). It's the innermost layer and has no dependencies on other layers.
Usecase: Implements the application's specific business logic (e.g., "Register User," "Submit Review"). It orchestrates data flow using the interfaces defined in the Domain layer.
Infrastructure: Provides concrete implementations for external concerns like the database (PostgresUserRepository, PostgresMovieRepository, etc.). It implements the interfaces defined in the Domain layer.
Delivery: Handles external interactions, such as HTTP requests. It receives requests, calls the appropriate Usecase, and formats responses.
Dependency Rule: Dependencies always flow inwards. For example, the Delivery layer depends on the Usecase layer, which depends on the Domain layer. The Infrastructure layer implements Domain interfaces, allowing for easy swapping of external services (like databases) without affecting core business logic.
# Getting Started
Follow these steps to get the Movie Review Backend API up and running on your local machine.
Prerequisites
Make sure you have the following installed:
Go: Version 1.20 or higher.
Download Go
PostgreSQL: A running PostgreSQL database server.
Download PostgreSQL
Git: For cloning the repository.
Download Git
Setup Instructions
Clone the Repository:
Open your terminal or command prompt and clone the project:
git clone https://github.com/your-username/movie-review-backend.git # Replace with your actual repo URL
cd movie-review-backend


Database Setup:
Create a Database: Connect to your PostgreSQL server (e.g., using psql or a GUI tool like DBeaver/pgAdmin) and create a new database.
CREATE DATABASE movie_reviews_db;


Table Creation: The application will automatically create the necessary tables (users, movies, reviews) with all validations and relationships when it starts up for the first time. You don't need to run separate migration scripts.
Environment Variables:
The application uses environment variables for database connection details and the HTTP port. Create a .env file in the root directory of the project (where main.go is located) and populate it like this:
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=movie_reviews_db
HTTP_PORT=8080


Replace your_db_user and your_db_password with your PostgreSQL credentials.
You can change DB_HOST, DB_PORT, DB_NAME, and HTTP_PORT if needed.
Run the Application:
Once your environment variables are set, you can start the backend server:
go run cmd/api/main.go

You should see output indicating that the database connection was successful and the server is starting on the specified HTTP port (e.g., Starting HTTP server on port 8080...).
# API Endpoints
The API is accessible via HTTP requests. Below are the main endpoints. All requests and responses are in JSON format.
Users
Register a New User
URL: /users/register
Method: POST
Description: Creates a new user account.
Request Body:
{
    "name": "John Doe",
    "email": "john.doe@example.com",
    "password": "securepassword123",
    "bio": "Movie enthusiast!",
    "gender": "Male",
    "profile_picture": "http://example.com/profile.jpg"
}


Success Response (201 Created):
{
    "name": "John Doe",
    "email": "john.doe@example.com",
    "bio": "Movie enthusiast!",
    "gender": "Male",
    "joined_on": "2024-07-30T14:00:00Z",
    "profile_picture": "http://example.com/profile.jpg"
}


Error Responses: 400 Bad Request (missing fields), 409 Conflict (email already exists), 500 Internal Server Error.
User Login
URL: /users/login
Method: POST
Description: Authenticates a user and returns a JWT token.
Request Body:
{
    "email": "john.doe@example.com",
    "password": "securepassword123"
}


Success Response (200 OK):
{
    "token": "your_generated_jwt_token_here"
}


Error Responses: 400 Bad Request, 401 Unauthorized (invalid credentials), 500 Internal Server Error.
Get User by Email
URL: /users/{email}
Method: GET
Description: Retrieves a user's profile details.
URL Parameters: email (e.g., john.doe@example.com)
Success Response (200 OK): Same as Register success response.
Error Responses: 404 Not Found, 500 Internal Server Error.
Update User Profile
URL: /users/{email}
Method: PUT
Description: Updates an existing user's profile information.
URL Parameters: email
Request Body: (Fields are optional; only provided fields will be updated)
{
    "name": "John D.",
    "bio": "Updated bio for movie lover",
    "password": "newsecurepassword"
}


Success Response (200 OK): Updated user object.
Error Responses: 400 Bad Request, 404 Not Found, 500 Internal Server Error.
Delete User
URL: /users/{email}
Method: DELETE
Description: Deletes a user account and all associated reviews.
URL Parameters: email
Success Response (204 No Content): No body.
Error Responses: 404 Not Found, 500 Internal Server Error.
Movies
Create a New Movie
URL: /movies
Method: POST
Description: Adds a new movie to the database.
Request Body:
{
    "movie_id": 1001,
    "title": "The Go Movie",
    "description": "A thrilling adventure about concurrency.",
    "release_on": "2024-01-15T00:00:00Z",
    "images": ["url1.jpg", "url2.jpg"],
    "videos": ["trailer.mp4"],
    "genres": ["Action", "Comedy"],
    "directors": ["Gopher Director"],
    "writes": ["Go Lang"],
    "casts": ["Goroutine", "Channel"],
    "average_ratings": 5,
    "origin_country": "USA",
    "languages": ["English"],
    "production_companies": ["Go Studios"],
    "budget": 100000000,
    "runtime": "2h 15m"
}


Note: movie_id is expected to be provided in the request body. In a real-world scenario, this might come from an external movie API (like TMDB) or be generated by your application logic.
Success Response (201 Created): Created movie object.
Error Responses: 400 Bad Request, 500 Internal Server Error.
Get Movie by ID
URL: /movies/{id}
Method: GET
Description: Retrieves details for a specific movie.
URL Parameters: id (e.g., 1001)
Success Response (200 OK): Movie object.
Error Responses: 404 Not Found, 500 Internal Server Error.
Update Movie Details
URL: /movies/{id}
Method: PUT
Description: Updates details for an existing movie.
URL Parameters: id
Request Body: (Fields are optional; only provided fields will be updated)
{
    "title": "The Go Movie: The Sequel",
    "average_ratings": 4
}


Success Response (200 OK): Updated movie object.
Error Responses: 400 Bad Request, 404 Not Found, 500 Internal Server Error.
Delete Movie
URL: /movies/{id}
Method: DELETE
Description: Deletes a movie and all its associated reviews.
URL Parameters: id
Success Response (204 No Content): No body.
Error Responses: 404 Not Found, 500 Internal Server Error.
Reviews
Create a New Review
URL: /reviews
Method: POST
Description: Submits a new review for a movie by a user.
Request Body:
{
    "review_id": 2001, // For demo, usually generated by DB
    "movie_id": 1001,
    "user_id": 1,      // This is the 'user_id' (int64) from the users table
    "title": "Amazing Movie!",
    "description": "Loved the plot and characters.",
    "rating": 5,
    "is_spoiler": false
}


Important: The user_id here must be the integer ID of an existing user, not their email.
Success Response (201 Created): Created review object.
Error Responses: 400 Bad Request, 404 Not Found (movie or user not found), 409 Conflict (user already reviewed this movie), 500 Internal Server Error.
Get Review by ID
URL: /reviews/{id}
Method: GET
Description: Retrieves details for a specific review.
URL Parameters: id (e.g., 2001)
Success Response (200 OK): Review object.
Error Responses: 404 Not Found, 500 Internal Server Error.
Get Reviews by Movie ID
URL: /reviews/movie/{movie_id}
Method: GET
Description: Retrieves all reviews for a given movie.
URL Parameters: movie_id (e.g., 1001)
Success Response (200 OK): Array of review objects.
Error Responses: 500 Internal Server Error.
Update Review
URL: /reviews/{id}
Method: PUT
Description: Updates an existing review.
URL Parameters: id
Request Body: (Fields are optional)
{
    "title": "Still Amazing!",
    "description": "Re-watched and still love it.",
    "rating": 5,
    "is_spoiler": true
}


Success Response (200 OK): Updated review object.
Error Responses: 400 Bad Request, 404 Not Found, 500 Internal Server Error.
Delete Review
URL: /reviews/{id}
Method: DELETE
Description: Deletes a specific review.
URL Parameters: id
Success Response (204 No Content): No body.
Error Responses: 404 Not Found, 500 Internal Server Error.
Like a Review
URL: /reviews/{id}/like
Method: POST
Description: Increments the like count for a review.
URL Parameters: id
Success Response (200 OK): No body (or could return updated likes count).
Error Responses: 404 Not Found, 500 Internal Server Error.
Dislike a Review
URL: /reviews/{id}/dislike
Method: POST
Description: Increments the dislike count for a review.
URL Parameters: id
Success Response (200 OK): No body (or could return updated dislikes count).
Error Responses: 404 Not Found, 500 Internal Server Error.
