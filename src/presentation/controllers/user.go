package controllers

import (
	"github.com/dscamargo/crud-clean-architecture/src/domain"
	"github.com/dscamargo/crud-clean-architecture/src/user/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	createUserUsecase usecase.CreateUserUsecase
	getUserUsecase    usecase.GetUserUsecase
}

func NewUserController(createUser usecase.CreateUserUsecase, getUser usecase.GetUserUsecase) UserController {
	return UserController{createUser, getUser}
}

func (controller UserController) CreateUser(c *gin.Context) {
	input := struct {
		Name                 string `json:"name"`
		Email                string `json:"email"`
		Password             string `json:"password"`
		PasswordConfirmation string `json:"passwordConfirmation"`
	}{}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(500, map[string]interface{}{"error": domain.ErrInternalServerError.Error()})
		return
	}
	userId, err := controller.createUserUsecase.CreateUser(input.Name, input.Email, input.Password, input.PasswordConfirmation)
	if err != nil {
		c.JSON(400, map[string]interface{}{"message": err.Error()})
		return
	}
	c.JSON(201, map[string]interface{}{"userId": userId})
	return
}

func (controller UserController) GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	user, err := controller.getUserUsecase.GetUserById(id)
	if err != nil {
		if err == domain.ErrNotFound {
			c.JSON(404, map[string]interface{}{"message": err.Error()})
			return
		}
		c.JSON(500, map[string]interface{}{"message": domain.ErrInternalServerError.Error()})
		return
	}
	c.JSON(200, user)
	return
}
