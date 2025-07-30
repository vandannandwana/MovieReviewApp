# **Movie Review Backend API**

This project is a robust backend API for a movie review platform, built with **Go** and designed following **Clean Architecture** principles. It provides core functionalities for user management, movie information, and user-submitted reviews.

## ** Features**

* **User Management**: Register new users, log in, retrieve, update, and delete user profiles.  
* **Movie Management**: Add new movies, fetch movie details by ID, update existing movie information, and delete movies.  
* **Review Management**: Submit new reviews for movies, retrieve individual reviews, get all reviews for a specific movie, update review content, and delete reviews.  
* **Review Interactions**: Like and dislike reviews.  
* **Data Integrity**: Enforces that each user can submit **only one review per movie** at the database level.  
* **PostgreSQL Database**: Uses PostgreSQL for persistent data storage.  
* **Clean Architecture**: Structured for high maintainability, testability, and scalability.

## ** Architecture Overview**

This application adheres to **Clean Architecture** principles, which organize code into distinct layers with strict dependency rules. This approach makes the codebase easier to understand, test, and evolve.

The main layers are:

* **Domain**: Contains the core business entities (User, Movie, Review) and their interfaces (e.g., UserRepository). It's the innermost layer and has no dependencies on other layers.  
* **Usecase**: Implements the application's specific business logic (e.g., "Register User," "Submit Review"). It orchestrates data flow using the interfaces defined in the Domain layer.  
* **Infrastructure**: Provides concrete implementations for external concerns like the database (PostgresUserRepository, PostgresMovieRepository, etc.). It implements the interfaces defined in the Domain layer.  
* **Delivery**: Handles external interactions, such as HTTP requests. It receives requests, calls the appropriate Usecase, and formats responses.

**Dependency Rule**: Dependencies always flow inwards. For example, the Delivery layer depends on the Usecase layer, which depends on the Domain layer. The Infrastructure layer implements Domain interfaces, allowing for easy swapping of external services (like databases) without affecting core business logic.

## ** Getting Started**

Follow these steps to get the Movie Review Backend API up and running on your local machine.

### **Prerequisites**

Make sure you have the following installed:

* **Go**: Version 1.20 or higher.  
  * [Download Go](https://go.dev/dl/)  
* **PostgreSQL**: A running PostgreSQL database server.  
  * [Download PostgreSQL](https://www.postgresql.org/download/)  
* **Git**: For cloning the repository.  
  * [Download Git](https://git-scm.com/downloads)

### **Setup Instructions**

1. Clone the Repository:  
   Open your terminal or command prompt and clone the project:  
   git clone https://github.com/vandannandwana/MovieReviewApp.git 
   cd movieReviewApp

2. **Database Setup**:  
   * **Create a Database**: Connect to your PostgreSQL server (e.g., using psql or a GUI tool like DBeaver/pgAdmin) and create a new database.  
     CREATE DATABASE postgres;

   * **Table Creation**: The application will automatically create the necessary tables (users, movies, reviews) with all validations and relationships when it starts up for the first time. You don't need to run separate migration scripts.  
3. Environment Variables:  
   The application uses a YAML file for configuration(you can do respective changes in /config/local.yaml):  
   DB\_HOST=localhost  
   DB\_PORT=5432  
   DB\_USER=your\_db\_user  
   DB\_PASSWORD=your\_db\_password  
   DB\_NAME=movie\_reviews\_db  
   HTTP\_PORT=8080

   * **Replace your\_db\_user and your\_db\_password** with your PostgreSQL credentials.  
   * You can change DB\_HOST, DB\_PORT, DB\_NAME, and HTTP\_PORT if needed.  
4. Run the Application:  
   Once your environment variables are set, you can start the backend server:  
   For BASH: CONFIG_PATH=config/config.yaml go run cmd/api/main.go
   For Powershell: $env:CONFIG_PATH="./config/local.yaml"
   then, go run .\cmd\api\main.go

   You should see output indicating that the database connection was successful and the server is starting on the specified HTTP port (e.g., Starting HTTP server on port 8080...).

## ** API Endpoints**

The API is accessible via HTTP requests. Below are the main endpoints. All requests and responses are in **JSON** format.

### **Users**

#### **Register a New User**

* **URL**: /users/register  
* **Method**: POST  
* **Description**: Creates a new user account.  
* **Request Body**:  
  {  
      "name": "Vandan Nandwana",  
      "email": "vandan@example.com",  
      "password": "12345",  
      "bio": "Movie enthusiast\!",  
      "gender": "Male",  
      "profile\_picture": "http://example.com/profile.jpg"  
  }

#### **User Login**

* **URL**: /users/login  
* **Method**: POST  
* **Description**: Authenticates a user and returns a JWT token.  
* **Request Body**:  
  {  
      "email": "vandan@example.com",  
      "password": "12345"  
  } 

### **Movies**

#### **Create a New Movie**

* **URL**: /movies  
* **Method**: POST  
* **Description**: Adds a new movie to the database.  
* **Request Body**:  
  {  
      "title": "The Go Movie",  
      "description": "A thrilling adventure about concurrency.",  
      "release\_on": "2024-01-15T00:00:00Z",  
      "images": \["url1.jpg", "url2.jpg"\],  
      "videos": \["trailer.mp4"\],  
      "genres": \["Action", "Comedy"\],  
      "directors": \["Gopher Director"\],  
      "writes": \["Go Lang"\],  
      "casts": \["Goroutine", "Channel"\],  
      "average\_ratings": 5,  
      "origin\_country": "USA",  
      "languages": \["English"\],  
      "production\_companies": \["Go Studios"\],  
      "budget": 100000000,  
      "runtime": "2h 15m"  
  }
 
#### **Get Movie by ID**

* **URL**: /movies/{id}  
* **Method**: GET  
* **Description**: Retrieves details for a specific movie.  
* **URL Parameters**: id (e.g., 1001\)  

#### **Update Movie Details**

* **URL**: /movies/{id}  
* **Method**: PUT  
* **Description**: Updates details for an existing movie.  
* **URL Parameters**: id  
* **Request Body**: (Fields are optional; only provided fields will be updated)  
  {  
      "title": "The Go Movie: The Sequel",  
      "description": "A thrilling adventure about concurrency.", 
      "email": "abc@gmail.com"
      "images": \["url1.jpg", "url2.jpg"\],  
      "videos": \["trailer.mp4"\],  
      "genres": \["Action", "Comedy"\],  
      "directors": \["Gopher Director"\],  
      "writes": \["Go Lang"\],  
      "casts": \["Goroutine", "Channel"\], 
      "origin\_country": "USA",  
      "languages": \["English"\],  
      "production\_companies": \["Go Studios"\],  
      "budget": 100000000,  
      "runtime": "2h 15m"  
  }


#### **Delete Movie**

* **URL**: /movies/{user_email}/{id}  
* **Method**: DELETE  
* **Description**: Deletes a movie and all its associated reviews.  
* **URL Parameters**: id  

### **Reviews**

#### **Create a New Review**

* **URL**: /reviews  
* **Method**: POST  
* **Description**: Submits a new review for a movie by a user.  
* **Request Body**:  
  {  
      "movie\_id": 1001,  
      "user\_email": "abc@gmail.com",     
      "title": "Amazing Movie\!",  
      "description": "Loved the plot and characters.",  
      "rating": 5,  
      "is\_spoiler": false  
  } 

#### **Get Review by ID**

* **URL**: /reviews/{id}  
* **Method**: GET  
* **Description**: Retrieves details for a specific review.  
* **URL Parameters**: id (e.g., 2001\)

#### **Get Reviews by Movie ID**

* **URL**: /reviews/movie/{movie\_id}  
* **Method**: GET  
* **Description**: Retrieves all reviews for a given movie.  
* **URL Parameters**: movie\_id (e.g., 1001\)

#### **Update Review**

* **URL**: /reviews/{id}  
* **Method**: PUT  
* **Description**: Updates an existing review.  
* **URL Parameters**: id  
* **Request Body**: (Fields are optional)  
  {  
      "movie\_id":1,
      "user_email":"abc@gmail.com"
      "title": "Still Amazing\!",  
      "description": "Re-watched and still love it.",  
      "rating": 5,  
      "is\_spoiler": true  
  }

#### **Delete Review**

* **URL**: /reviews/{user_email}/{id}  
* **Method**: DELETE  
* **Description**: Deletes a specific review.  
* **URL Parameters**: id

#### **Get Review By Movie Id**

* **URL**: /reviews/movie/{id}  
* **Method**: DELETE  
* **Description**: Deletes a specific review.  
* **URL Parameters**: id

#### **Get Review By Email Id**

* **URL**: /reviews/email/{id}  
* **Method**: DELETE  
* **Description**: Deletes a specific review.  
* **URL Parameters**: id