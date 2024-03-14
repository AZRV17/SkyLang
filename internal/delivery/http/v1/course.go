package v1

import (
	"github.com/AZRV17/Skylang/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) initCourseRoutes(r *gin.Engine) {
	courses := r.Group("/courses")
	{
		courses.POST("/", h.createCourse)
		courses.GET("/", h.getAllCourses)
		courses.GET("/:id", h.getCourseById)
		courses.PUT("/:id", h.updateCourse)
		courses.DELETE("/:id", h.deleteCourse)
	}
}

func (h *Handler) createCourse(c *gin.Context) {
	var input service.CreateCourseInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, err := h.service.CourseService.CreateCourse(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, course)
}

func (h *Handler) getAllCourses(c *gin.Context) {
	courses, err := h.service.CourseService.GetAllCourses()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, courses)
}

func (h *Handler) getCourseById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, err := h.service.CourseService.GetCourseByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, course)
}

func (h *Handler) updateCourse(c *gin.Context) {
	var input service.UpdateCourseInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, err := h.service.CourseService.UpdateCourse(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, course)
}

func (h *Handler) deleteCourse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.CourseService.DeleteCourse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted"})
}

func (h *Handler) getCourseByTitle(c *gin.Context) {
	title := c.Query("title")
	course, err := h.service.CourseService.GetCourseByTitle(title)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, course)
}

func (h *Handler) filterCoursesByTitle(c *gin.Context) {
	title := c.Query("title")
	courses, err := h.service.CourseService.FilterCoursesByTitle(title)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, courses)
}

func (h *Handler) sortCourseByDate(c *gin.Context) {
	courses, err := h.service.CourseService.SortCourseByDate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, courses)
}

func (h *Handler) sortCourseByRating(c *gin.Context) {
	courses, err := h.service.CourseService.SortCourseByRating()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, courses)
}

func (h *Handler) sortCourseByTitle(c *gin.Context) {
	courses, err := h.service.CourseService.SortCourseByTitle()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, courses)
}