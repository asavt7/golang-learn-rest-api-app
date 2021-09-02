package handlers

import (
	"errors"
	"github.com/asavt7/todo/pkg/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createList(ctx *gin.Context) {
	id, err := getUserId(ctx)
	if err != nil {
		return
	}

	var input domain.TodoList
	err = ctx.BindJSON(&input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	idList, err := h.service.TodoList.Create(id, input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]int{
		"id": idList,
	})
	return
}

func getUserId(ctx *gin.Context) (int, error) {
	id, ok := ctx.Get(UserIdCtx)
	if !ok {
		NewErrorResponse(ctx, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		NewErrorResponse(ctx, http.StatusInternalServerError, "user id if of invalid type")
		return 0, errors.New("user id if of invalid type")
	}
	return idInt, nil
}

type getAllListsResponse struct {
	Data []domain.TodoList `json:"data"`
}

func (h *Handler) getAllLists(ctx *gin.Context) {
	id, err := getUserId(ctx)
	if err != nil {
		return
	}

	lists, err := h.service.TodoList.GetAllLists(id)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getAllListsResponse{Data: lists})

}

func (h *Handler) getListById(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	list, err := h.service.TodoList.GetListById(userId, listId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, list)

}

func (h *Handler) updateList(ctx *gin.Context) {

}

func (h *Handler) deleteList(ctx *gin.Context) {

	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.TodoList.Delete(userId, listId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.Status(http.StatusOK)
}
