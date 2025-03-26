package models

import "time"

type Post struct {
    ID        string    `gorm:"primaryKey"`
    Title     string
    Content   string
    AuthorID  string
    CreatedAt time.Time
}
