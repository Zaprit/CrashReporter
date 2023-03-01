package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var ValidPlatforms = map[string]struct{}{
	"PlayStation 3":                        {},
	"PlayStation Vita":                     {},
	"RPCS3":                                {},
	"Web: Safari":                          {},
	"Web: Google Chrome or Chromium Based": {},
	"Web: Firefox":                         {},
	"Web: Other":                           {},
}

type Report struct {
	// ID is for indexing, UUID is so that people can't sequentially guess reports
	ID               uint `gorm:"primarykey"`
	UUID             string
	Title            string
	Username         string
	UserID           uint
	Avatar           string `gorm:"-"`
	Type             string
	Platform         string
	Description      string
	Resolved         bool
	Evidence         bool
	Read             bool
	SubmitterIP      string
	SubmitTime       time.Time `gorm:"autoCreateTime"`
	Comments         []Comment
	Priority         string
	DiscordMessageID string
}

func (r *Report) BeforeCreate(tx *gorm.DB) error {
	newUuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	r.UUID = newUuid.String()
	return nil
}
