package handlers

import (
	"github.com/asavt7/todo/pkg/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createItem(ctx *gin.Context) {

	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	var input domain.TodoItem
	err = ctx.BindJSON(&input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	idList, err := h.service.TodoItem.Create(userId, listId, input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]int{
		"id": idList,
	})
	return

}

func (h *Handler) getAllItems(ctx *gin.Context) {

	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	items, err := h.service.TodoItem.GetAllItems(userId, listId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string][]domain.TodoItem{
		"data": items,
	})
	return

}

func (h *Handler) getItemById(ctx *gin.Context) {

}

func (h *Handler) updateItem(ctx *gin.Context) {

}

func (h *Handler) deleteItem(ctx *gin.Context) {

}
