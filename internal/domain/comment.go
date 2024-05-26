package domain

import "time"

type Comment struct {
	ID        int       `json:"id"`
	CourseID  int       `json:"course_id"`
	Course    *Course   `json:"-" gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE"`
	AuthorID  int       `json:"author_id"`
	Author    *User     `json:"author" gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
