package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/mr687/privy-be-test-go/dto"
	"gitlab.com/mr687/privy-be-test-go/service"
)

// Define contract for auth controller
type AuthController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JwtService
}

// NewAuthController factory for auth controller
func NewAuthController(authService service.AuthService, jwtService service.JwtService) AuthController {
	return &authController{authService, jwtService}
}

func (controller *authController) Register(ctx *gin.Context) {
	var registerData dto.RegisterDTO
	errDto := ctx.ShouldBind(&registerData)
	if errDto != nil {
		response := MakeErrorResponse("Failed to validate request", errDto.Error(), EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Check if user email already exists
	if controller.authService.IsDuplicateEmail(registerData.Email) {
		response := MakeErrorResponse("Email already exists", "Duplicate Email", EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}

	newUser, err := controller.authService.CreateUser(registerData)
	if err != nil {
		response := MakeErrorResponse("Failed to create user", err.Error(), EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	// Sign auth token
	token := controller.jwtService.GenerateToken(strconv.FormatUint(newUser.Id, 10))
	newUser.Token = token

	response := MakeResponse(true, "User created successfully", newUser)
	ctx.JSON(http.StatusCreated, response)
}

func (c *authController) Login(ctx *gin.Context) {
	var loginData dto.LoginDTO
	errDto := ctx.ShouldBind(&loginData)
	if errDto != nil {
		response := MakeErrorResponse("Failed to validate request", errDto.Error(), EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user, err := c.authService.ValidateLogin(loginData)
	if err != nil {
		response := MakeErrorResponse("Failed to login", err.Error(), EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	if user == nil {
		response := MakeErrorResponse("Failed to login", "Password mismatch", EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	// Sign auth token
	token := c.jwtService.GenerateToken(strconv.FormatUint(user.Id, 10))
	user.Token = token

	response := MakeResponse(true, "User logged in successfully", user)
	ctx.JSON(http.StatusOK, response)
}
