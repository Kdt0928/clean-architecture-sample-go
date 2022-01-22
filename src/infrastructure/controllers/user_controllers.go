package controllers

import (
	"clean-architecture-go-sample/entities"
	"clean-architecture-go-sample/interfaces/database"
	"clean-architecture-go-sample/usecase"
	"strconv"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) {
	u := entities.User{}
	c.Bind(&u)
	err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, nil)
}

func (controller *UserController) Index(c Context) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, users)
}

func (controller *UserController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, user)
}
