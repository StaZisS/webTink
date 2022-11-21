package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	listing "web"
)

func (h *Handler) signUp(c *gin.Context) {
	var input listing.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input listing.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, refreshToken, err := h.services.Authorization.GenerateToken(input.Password, input.Email)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.SetCookie("access_token", accessToken, viper.GetInt("ACCESS_TOKEN_MAXAGE")*60, "/", "localhost", false, true)
	c.SetCookie("refresh_token", refreshToken, viper.GetInt("REFRESH_TOKEN_MAXAGE")*60, "/", "localhost", false, true)
	c.SetCookie("logged_in", "true", viper.GetInt("ACCESS_TOKEN_MAXAGE")*60, "/", "localhost", false, false)

	c.JSON(http.StatusOK, map[string]interface{}{
		"token":         accessToken,
		"refresh_token": refreshToken,
	})
}
func (h *Handler) logOut(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	c.SetCookie("logged_in", "false", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}
func (h *Handler) refresh(c *gin.Context) {
	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	accessToken, err := h.services.Authorization.Refresh(cookie)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	c.SetCookie("access_token", accessToken, viper.GetInt("ACCESS_TOKEN_MAXAGE")*60, "/", "localhost", false, true)
	c.SetCookie("logged_in", "true", viper.GetInt("ACCESS_TOKEN_MAXAGE")*60, "/", "localhost", false, false)

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": accessToken,
	})
}
