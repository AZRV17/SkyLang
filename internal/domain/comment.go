package domain

import "time"

type Comment struct {
	ID        int       `json:"id"`
	CourseID  int       `json:"course_id"`
	Course    *Course   `json:"-" gorm:"foreignKey:CourseID"`
	AuthorID  int       `json:"author_id"`
	Author    *User     `json:"-" gorm:"foreignKey:AuthorID"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
