package v1

import (
	"github.com/AZRV17/Skylang/internal/service"
	"github.com/gin-gonic/gin"
	"log"
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
		courses.PUT("/:id/updateIcon", h.updateIcon)
		courses.GET("/:id/icon", h.getIcon)
		courses.PUT("/:id/updateGrate", h.updateCourseGrate)
		courses.GET("/author/:id", h.getCourseByAuthorId)
	}
}

type createCourseInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Language    string `json:"language"`
	Icon        string `json:"icon"`
	Author      int    `json:"author"`
}

func (h *Handler) createCourse(c *gin.Context) {
	var input createCourseInput
	if err := c.BindJSON(&input); err != nil {
		log.Println(input)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(input)

	icon := input.Icon

	input.Icon = ""

	course, err := h.service.CourseService.CreateCourse(service.CreateCourseInput{
		Name:        input.Name,
		Description: input.Description,
		Language:    input.Language,
		Icon:        icon,
		Author:      input.Author,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if icon == "" {
		c.JSON(http.StatusOK, course)
		return
	}

	err = h.service.ImageService.SetCourseImage(course.ID, icon)
	if err != nil {
		log.Println(err.Error())
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

type updateIconInput struct {
	Icon string `json:"icon"`
}

func (h *Handler) updateIcon(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input updateIconInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.ImageService.SetCourseImage(id, input.Icon)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Icon updated"})
}

func (h *Handler) getIcon(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	icon, err := h.service.ImageService.GetCourseIcon(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.File(icon.Name())
}

type updateGrateInput struct {
	UserID int `json:"user_id"`
	Grate  int `json:"grate"`
}

func (h *Handler) updateCourseGrate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input updateGrateInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grate := service.CreateRatingInput{
		CourseID: id,
		UserID:   input.UserID,
		Grate:    input.Grate,
	}

	err = h.service.CourseService.UpdateCourseGrate(id, &grate)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Grate updated"})
}

func (h *Handler) getCourseByAuthorId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	courses, err := h.service.CourseService.GetCourseByAuthorID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, courses)
}
