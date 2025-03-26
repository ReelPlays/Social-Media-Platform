package models

import "time"

type Comment struct {
    ID        string    `gorm:"primaryKey"`
    PostID    string
    Content   string
    AuthorID  string
    CreatedAt time.Time
}
