package domain

import "time"

type Lecture struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CourseID    int       `json:"course_id"`
	Course      *Course   `json:"course" gorm:"foreignKey:CourseID"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
