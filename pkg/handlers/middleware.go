package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const AuthorizationHeader = "Authorization"
const UserIdCtx = "UserId"

func (h *Handler) userIdentity(c *gin.Context) {

	header := c.GetHeader(AuthorizationHeader)
	if header == "" {
		NewErrorResponse(c, http.StatusBadRequest, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	userId, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(UserIdCtx, userId)
}
