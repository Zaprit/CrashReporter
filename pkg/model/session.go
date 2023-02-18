package model

import (
    "github.com/google/uuid"
    "gorm.io/gorm"
    "time"
)

type Session struct {
    ID string `gorm:"primarykey" json:"-"`
    Username string `json:"login"`
    OAuthToken string `json:"-"`
    AvatarURL string `json:"avatar_url"`
    Expires time.Time `json:"-"`
}

func (s *Session) BeforeCreate(tx *gorm.DB) error {
    s.Expires = time.Now().UTC().AddDate(0,0,7)

    newUuid, err := uuid.NewRandom()
    if err != nil {
        return err
    }
    s.ID = newUuid.String()
    return nil
}