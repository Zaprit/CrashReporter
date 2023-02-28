package model

import (
	"encoding/json"
	"time"
)

type Comment struct {
	ID uint `gorm:"primarykey" json:"-"`

	// Username of whoever posted it, keeping it simple here
	Poster       string `json:"poster"`
	PosterAvatar string `json:"poster_avatar"`

	ReportID uint `json:"-"`

	CreateTime time.Time `json:"-"`

	// The actual body of the comment
	Content string `json:"content"`
}

func (d *Comment) MarshalJSON() ([]byte, error) {
	type Alias Comment
	return json.Marshal(&struct {
		*Alias
		CreateTime string `json:"create_time"`
	}{
		Alias:      (*Alias)(d),
		CreateTime: d.CreateTime.Format("Monday January 02 2006, 3:04:05 PM MST"),
	})
}
