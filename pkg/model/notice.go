package model

import "time"

type Notice struct {
	ID         uint      `gorm:"primarykey"`
	Status     string    `json:"status"`
	NoticeText string    `json:"notice_text"`
	Date       time.Time `json:"date"`
	Ended      bool      `json:"-"`
}
