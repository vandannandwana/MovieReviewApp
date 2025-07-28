package http

import (
	"github.com/gin-gonic/gin"
	"github.com/vandannandwana/MovieReviewApp/internal/delivery/http/handler"
)

func NewRouter(userHandler handler.UserHandler, movieHandler handler.MovieHandler, reviewHandler handler.ReviewHandler) (*gin.Engine){

	router := gin.Default()

	//User-Routes
	userRouter := router.Group("/users")
	userRouter.POST("/register", userHandler.RegisterUser)
	userRouter.POST("/login", userHandler.LoginUser)

	movieRouter := router.Group("/movies")

	movieRouter.POST("/", movieHandler.CreateMovie)
	movieRouter.GET("/{id}", movieHandler.GetMovieById)
	movieRouter.PUT("/{id}", movieHandler.UpdateMovie)
	movieRouter.DELETE("/{id}", movieHandler.DeleteMovie)

	reviewRouter := router.Group("/reviews")

	reviewRouter.POST("/", reviewHandler.CreateReview)
	reviewRouter.GET("/{id}", reviewHandler.GetReviewById)
	reviewRouter.GET("/movie/{id}", reviewHandler.GetReviewByMovieId)
	reviewRouter.GET("/email/{id}", reviewHandler.GetReviewByUserEmailId)
	reviewRouter.PUT("/", reviewHandler.UpdateReview)
	reviewRouter.DELETE("/", reviewHandler.DeleteReview)

	return router
}

