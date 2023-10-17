package models

import "time"

type Post struct {
	Id        int64  `gorm:"primaryKey" json:"id"`
	Title     string `gorm:"type:varchar(255)" json:"title"`
	Content   string `gorm:"type:text" json:"content"`
	UserId    int64  `json:"user_id"`
	User      User   `gorm:"foreignKey:UserId" json:"-"` // Relasi Many-to-One
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Username string `gorm:"type:varchar(255)" json:"username"`
	Password string `gorm:"type:varchar(255)" json:"-"`
	Posts    []Post `json:"posts"` // Relasi One-to-Many
}
