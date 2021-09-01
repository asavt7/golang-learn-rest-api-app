package handlers

import (
	"github.com/asavt7/todo/pkg/domain"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var input domain.User

	err := ctx.BindJSON(&input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := h.service.CreateUser(input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]int{
		"id": userId,
	})

	logrus.Info(userId)
}

func (h *Handler) signIn(ctx *gin.Context) {

}
