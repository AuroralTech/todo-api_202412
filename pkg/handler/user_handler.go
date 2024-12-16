package handler

import (
	"log"
	"net/http"

	"github.com/AuroralTech/todo-api_202412/pkg/usecase"

	"github.com/labstack/echo"
)

type UserHandler interface {
	UpdateUser(ctx echo.Context) error
}

type updateUser struct {
	updateUserUsecase usecase.UserUsecase
}

type UpdateUserRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewUserHandler(updateUserUsecase usecase.UserUsecase) UserHandler {
	return &updateUser{updateUserUsecase: updateUserUsecase}
}

func (h *updateUser) UpdateUser(ctx echo.Context) error {
	req := UpdateUserRequest{}
	if err := ctx.Bind(&req); err != nil {
		log.Printf("bind error occurred in user handler: %v", err)
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}
	updatedCounts, err := h.updateUserUsecase.UpdateUser(ctx.Request().Context(), usecase.UserUsecaseInput{
		ID:   req.ID,
		Name: req.Name,
	})
	if err != nil {
		log.Printf("error occurred in user handler: %v", err)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, echo.Map{
		"updated_counts": updatedCounts,
	})
}
