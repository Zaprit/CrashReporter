package model

import "time"

type Comment struct {
	ID uint `gorm:"primarykey" json:"-"`

	// Username of whoever posted it, keeping it simple here
	Poster       string `json:"poster"`
	PosterAvatar string `json:"poster_avatar"`

	ReportID uint `json:"-"`

	CreateTime time.Time `json:"create_time"`

	// The actual body of the comment
	Content string `json:"content"`
}
