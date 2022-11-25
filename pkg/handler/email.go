package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	listing "web"
)

func (h *Handler) sendMail(c *gin.Context) {
	var input listing.Email
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Email.SendEmail(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
