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
    ID string `gorm:"primarykey"`
    Title string
    Username string
    Type string
    Platform string
    Description string
    Evidence bool
}

// BeforeCreate will set a UUID rather than numeric ID.
func (r *Report) BeforeCreate(tx *gorm.DB) error {
    newUuid, err := uuid.NewRandom()
    if err != nil {
        return err
    }
    r.ID = newUuid.String()
    return nil
}