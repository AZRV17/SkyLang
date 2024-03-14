package v1

import (
	"github.com/AZRV17/Skylang/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) initLectureRoutes(r *gin.Engine) {
	lectures := r.Group("/lectures")
	{
		lectures.POST("/", h.createLecture)
		lectures.GET("/", h.getAllLectures)
		lectures.GET("/:id", h.getLectureById)
		lectures.PUT("/:id", h.updateLecture)
		lectures.DELETE("/:id", h.deleteLecture)
	}
}

func (h *Handler) createLecture(c *gin.Context) {
	var input service.CreateLectureInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lecture, err := h.service.LectureService.CreateLecture(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lecture)
}

func (h *Handler) getAllLectures(c *gin.Context) {
	lectures, err := h.service.LectureService.GetAllLectures()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lectures)
}

func (h *Handler) getLectureById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lecture, err := h.service.LectureService.GetLectureByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lecture)
}

func (h *Handler) updateLecture(c *gin.Context) {
	var input service.UpdateLectureInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lecture, err := h.service.LectureService.UpdateLecture(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lecture)
}

func (h *Handler) deleteLecture(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.LectureService.DeleteLecture(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lecture deleted"})
}
