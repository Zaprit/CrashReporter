package model

import "time"

type Comment struct {
    ID uint `gorm:"primarykey" json:"-"`

    // Username of whoever posted it, keeping it simple here
    Poster string `json:"poster"`

    ReportID uint `json:"-"`

    CreateTime time.Time `json:"create_time"`
}