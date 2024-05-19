package v1

import (
	"github.com/AZRV17/Skylang/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) initExerciseRoutes(r *gin.Engine) {
	exercises := r.Group("/exercises")
	{
		exercises.POST("/", h.createExercise)
		exercises.GET("/", h.getAllExercises)
		exercises.GET("/:id", h.getExerciseById)
		exercises.PUT("/:id", h.updateExercise)
		exercises.DELETE("/:id", h.deleteExercise)
		exercises.GET("/course/:id", h.getExercisesByCourse)
	}
}

func (h *Handler) createExercise(c *gin.Context) {
	var input service.CreateExerciseInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exercise, err := h.service.ExerciseService.CreateExercise(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exercise)
}

func (h *Handler) getAllExercises(c *gin.Context) {
	exercises, err := h.service.ExerciseService.GetAllExercises()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exercises)
}

func (h *Handler) getExerciseById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exercise, err := h.service.ExerciseService.GetExerciseByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exercise)
}

func (h *Handler) updateExercise(c *gin.Context) {
	var input service.UpdateExerciseInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exercise, err := h.service.ExerciseService.UpdateExercise(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exercise)
}

func (h *Handler) deleteExercise(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.ExerciseService.DeleteExercise(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Exercise deleted"})
}

func (h *Handler) getExercisesByCourse(c *gin.Context) {
	courseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exercises, err := h.service.ExerciseService.GetExercisesByCourseID(courseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exercises)
}
