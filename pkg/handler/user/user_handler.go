package user

import (
	"log"
	"net/http"

	usecase "github.com/AuroralTech/todo-api_202412/pkg/usecase/user"
	"github.com/labstack/echo"
)

type UpdateUserRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(userUsecase *usecase.UserUsecase) UserHandler {
	return UserHandler{userUsecase: userUsecase}
}

func (h *UserHandler) UpdateUser(ctx echo.Context) error {
	req := UpdateUserRequest{}
	if err := ctx.Bind(&req); err != nil {
		log.Printf("bind error occurred in user handler: %v", err)
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}
	updatedCounts, err := h.userUsecase.Execute(ctx.Request().Context(), usecase.UpdateUserInput{
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
