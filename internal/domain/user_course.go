package domain

import (
	"time"
)

type UserCourse struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"user_id"`
	User      User      `json:"-" gorm:"foreignKey:UserID"`
	CourseID  int       `json:"course_id"`
	Course    Course    `json:"-" gorm:"foreignKey:CourseID"`
	Status    string    `json:"status" gorm:"default:in_progress"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
