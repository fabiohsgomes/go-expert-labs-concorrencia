package user_controller

import (
	"context"
	"fullcycle-auction_go/configuration/rest_err"
	"fullcycle-auction_go/internal/infra/api/web/validation"
	"fullcycle-auction_go/internal/usecase/user_usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	userUseCase user_usecase.UserUseCaseInterface
}

func NewUserController(userUseCase user_usecase.UserUseCaseInterface) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var userInput user_usecase.UserInputDTO

	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		restErr := validation.ValidateErr(err)

		ctx.JSON(restErr.Code, restErr)
		return
	}

	userData, err := c.userUseCase.CreateUser(context.Background(), userInput)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		ctx.JSON(errRest.Code, errRest)
		return
	}

	ctx.JSON(http.StatusCreated, userData)
}

func (u *UserController) FindUsers(c *gin.Context) {
	userData, err := u.userUseCase.FindUsers(context.Background())
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, userData)
}

func (u *UserController) FindUserById(c *gin.Context) {
	userId := c.Param("userId")

	if err := uuid.Validate(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid fields", rest_err.Causes{
			Field:   "userId",
			Message: "Invalid UUID value",
		})

		c.JSON(errRest.Code, errRest)
		return
	}

	userData, err := u.userUseCase.FindUserById(context.Background(), userId)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, userData)
}
