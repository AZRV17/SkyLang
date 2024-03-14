package domain

import "time"

type Course struct {
	ID          int        `json:"id" gorm:"primaryKey"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Language    string     `json:"language"`
	Icon        string     `json:"icon"`
	Grate       float32    `json:"grate"`
	Exercises   []Exercise `json:"exercises" gorm:"foreignKey:CourseID"`
	Lectures    []Lecture  `json:"lectures" gorm:"foreignKey:CourseID"`
	AuthorID    int        `json:"author_id"`
	Author      *User      `json:"-" gorm:"foreignKey:AuthorID"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`
}
