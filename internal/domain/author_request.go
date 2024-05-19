package domain

type AuthorRequest struct {
	ID     int   `json:"id" gorm:"primaryKey"`
	UserID int   `json:"user_id"`
	User   *User `json:"user" gorm:"foreignKey:UserID"`
}
