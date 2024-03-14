package v1

import (
	"errors"
	"github.com/AZRV17/Skylang/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) initUserRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("/:id", h.getUserById)
		users.GET("/", h.getAllUsers)
		users.POST("/signup", h.signUp)
		users.POST("/login", h.singInByLogin)
		users.POST("/login_email", h.signInByEmail)
		users.PUT("/:id", h.updateUser)
		users.DELETE("/:id", h.deleteUser)
		users.PUT("/:id/changePassword", h.updatePassword)
		users.PUT("/:id/signUpForCourse", h.signUpForCourse)
	}
}

func (h *Handler) getUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.UserService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.service.UserService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

type SignInByLoginInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *Handler) singInByLogin(c *gin.Context) {
	var input SignInByLoginInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: return JWT token
	user, err := h.service.UserService.SignInByLogin(input.Login, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

type SignInByEmailInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *Handler) signInByEmail(c *gin.Context) {
	var input SignInByEmailInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.UserService.SignInByEmail(input.Login, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

type SignUpInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email" binding:"required,email"`
	Role     string `json:"role"`
}

func (h *Handler) signUp(c *gin.Context) {
	var input SignUpInput

	if err := c.BindJSON(&input); err != nil {
		err = errors.New("invalid input: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.UserService.SignUp(service.CreateUserInput(input))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) updateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input service.UpdateUserInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateUserInput := input
	updateUserInput.ID = id

	user, err := h.service.UserService.UpdateUser(updateUserInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

type UpdatePasswordInput struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
}

func (h *Handler) updatePassword(c *gin.Context) {
	var input UpdatePasswordInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.UserService.UpdatePassword(input.ID, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) deleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.UserService.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

type signUpForCourseInput struct {
	UserID   int `json:"user_id"`
	CourseID int `json:"course_id"`
}

func (h *Handler) signUpForCourse(c *gin.Context) {
	var input signUpForCourseInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.UserService.SignUpForCourse(input.UserID, input.CourseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User signed up for course"})
}
