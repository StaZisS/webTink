package handler

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
	"strings"
	listing "web"
)

type getAllPostsResponse struct {
	Data []listing.Post `json:"data"`
}

func (h *Handler) getAllPosts(c *gin.Context) {
	lists, err := h.services.Post.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	for i := range lists {
		file, err := os.Open("../images/" + lists[i].ID.String() + ".txt")
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		defer file.Close()
		reader := bufio.NewReader(file)
		lists[i].Photo, err = reader.ReadString('\n')
		lists[i].Photo = strings.Trim(lists[i].Photo, "\n")
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.JSON(http.StatusOK, getAllPostsResponse{
		Data: lists,
	})
}
func (h *Handler) getPostById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid post id param")
		return
	}
	list, err := h.services.Post.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	file, err := os.Open("../images/" + list.ID.String() + ".txt")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	list.Photo, err = reader.ReadString('\n')
	list.Photo = strings.Trim(list.Photo, "\n")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}
func (h *Handler) createPost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input listing.Post
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Post.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	file, err := os.Create("../images/" + id.String() + ".txt")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()
	file.WriteString(input.Photo + "\n")
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) updatePost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid post id param")
		return
	}
	var input listing.UpdatePostInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.Post.Update(userId, id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := os.Remove("../images/" + id.String() + ".txt"); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	file, err := os.Create("../images/" + id.String() + ".txt")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()
	file.WriteString(input.Photo + "\n")
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "ok",
	})
}
func (h *Handler) deletePost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid post id param")
		return
	}
	if err := os.Remove("../images/" + id.String() + ".txt"); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.services.Post.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "ok",
	})
}
