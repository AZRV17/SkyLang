package domain

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        int          `json:"id" gorm:"primaryKey"`
	Login     string       `json:"login"`
	Password  string       `json:"password"`
	Email     string       `json:"email"`
	Role      string       `json:"role"`
	Avatar    string       `json:"avatar,omitempty" gorm:"null"`
	Courses   []UserCourse `json:"courses"`
	CreatedAt time.Time    `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time    `json:"updatedAt" gorm:"autoUpdateTime"`
}

func (u User) CheckPassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return &ErrInvalidPassword{}
	}

	return nil
}
