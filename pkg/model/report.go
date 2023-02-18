package model

import "gorm.io/gorm"
import "github.com/google/uuid"

var ValidPlatforms = map[string]struct{} {
    "PlayStation 3":{},
    "PlayStation Vita": {},
    "RPCS3": {},
    "Web: Safari": {},
    "Web: Google Chrome or Chromium Based": {},
    "Web: Firefox": {},
    "Web: Other": {},
}

type Report struct {
    ID uint `gorm:"primarykey"`
    UUID string
    Title string
    Username string
    Type string
    Platform string
    Description string
    Evidence bool
}

func (r *Report) BeforeCreate(tx *gorm.DB) error {
    newUuid, err := uuid.NewRandom()
    if err != nil {
        return err
    }
    r.UUID = newUuid.String()
    return nil
}