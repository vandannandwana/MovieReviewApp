package http

import (
	"github.com/gin-gonic/gin"
	"github.com/vandannandwana/MovieReviewApp/internal/delivery/http/handler"
)

func NewRouter(userHandler handler.UserHandler) (*gin.Engine){

	router := gin.Default()

	//User-Routes
	userRouter := router.Group("/users")
	userRouter.POST("/register", userHandler.RegisterUser)
	userRouter.POST("/login", userHandler.LoginUser)

	// movieRouter := router.Group("/movies")
	// reviewRouter := router.Group("/reviews")

	return router
}

