package v1

import (
	"github.com/AZRV17/Skylang/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func (h *Handler) initAuthorRequestRoutes(r *gin.Engine) {
	requests := r.Group("/requests")
	{
		requests.GET("/", h.getAllAuthorRequests)
		requests.GET("/:id", h.getAuthorRequestById)
		requests.POST("/", h.createAuthorRequest)
		requests.DELETE("/:id", h.deleteAuthorRequest)
		requests.GET("/user/:id", h.getAuthorRequestByUserID)
	}
}

func (h *Handler) getAllAuthorRequests(c *gin.Context) {
	authorRequests, err := h.service.AuthorRequestsService.GetAuthorRequests()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, authorRequests)
}

func (h *Handler) getAuthorRequestById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authorRequest, err := h.service.AuthorRequestsService.GetAuthorRequestByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, authorRequest)
}

type CreateAuthorRequestInput struct {
	UserID int `json:"author_id" binding:"required"`
}

func (h *Handler) createAuthorRequest(c *gin.Context) {
	var input CreateAuthorRequestInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authorRequest := &service.CreateAuthorRequestInput{
		UserID: input.UserID,
	}

	_, err := h.service.AuthorRequestsService.CreateAuthorRequest(*authorRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Author request created"})
}

func (h *Handler) deleteAuthorRequest(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.AuthorRequestsService.DeleteAuthorRequest(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Author request deleted"})
}

func (h *Handler) getAuthorRequestByUserID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authorRequest, err := h.service.AuthorRequestsService.GetAuthorRequestByUserID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if authorRequest == nil || authorRequest.User == nil && authorRequest.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author request not found"})
		return
	}

	c.JSON(http.StatusOK, authorRequest)
}
