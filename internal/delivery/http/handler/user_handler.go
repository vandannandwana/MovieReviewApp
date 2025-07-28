package handler

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/vandannandwana/MovieReviewApp/internal/delivery/http/dto"
	"github.com/vandannandwana/MovieReviewApp/internal/usecase"
)

type UserHandler struct {
	userService usecase.UserService
}

func NewUserHandler (userService usecase.UserService) *UserHandler{
	return &UserHandler{userService : userService}
}

func (h *UserHandler) RegisterUser(c *gin.Context){

	var req dto.RegisterUserRequest

	if err := c.ShouldBind(&req); err !=nil{
		c.JSON(http.StatusBadRequest, "Invalid request payload")
		return 
	}

	if req.Name == "" || req.Email == "" || req.Password == ""{
		c.JSON(http.StatusBadRequest, "Name, email, and password are required")
		return
	}

	user, err := h.userService.RegisterUser(req.Name, req.Email, req.Password, req.Bio, req.Gender, req.ProfilePicture)

	if err != nil{
		
		c.JSON(http.StatusBadRequest, "failed to register user: "+err.Error())
		return

	}

	c.JSON(http.StatusCreated, dto.UserResponse{
		Name:           user.Name,
        Email:          user.Email,
        Bio:            user.Bio,
        Gender:         user.Gender,
        JoinedOn:       user.JoinedOn,
        ProfilePicture: user.ProfilePicture,
	})

}

func (h *UserHandler) LoginUser(c *gin.Context){
	var req dto.LoginUserRequest

	if err := c.ShouldBind(&req); err != nil{
		c.JSON(http.StatusBadRequest, "invalid request payloads")
		return
	}

	if req.Email == "" || req.Password == "" {
        c.JSON(http.StatusBadRequest, "Email and password are required")
        return
    }

	isLoggedIn, err := h.userService.LoginUser(req.Email, req.Password)

	if err !=nil{
		c.JSON(http.StatusInternalServerError, "Failed to login: "+err.Error())
		return
	}

	if isLoggedIn{
		c.JSON(http.StatusAccepted, dto.LoginResponse{Token: fmt.Sprint(isLoggedIn)})
	}else {	
		c.JSON(http.StatusNotAcceptable, dto.LoginResponse{Token: fmt.Sprint(isLoggedIn)})
	}

}
