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
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(ctx.Param("item_id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	item, err := h.service.TodoItem.GetById(userId, itemId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, item)
	return
}

func (h *Handler) updateItem(ctx *gin.Context) {

	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(ctx.Param("item_id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var input domain.UpdateTodoItem
	err = ctx.BindJSON(&input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err = input.Validate()
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.TodoItem.Update(userId, itemId, input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
	return

}

func (h *Handler) deleteItem(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(ctx.Param("item_id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.TodoItem.Delete(userId, itemId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
	return
}
