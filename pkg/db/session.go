package db

import (
	"errors"
	"github.com/Zaprit/CrashReporter/pkg/model"
	"time"
)

func SaveSession(session *model.Session) error {
	database.Where("username = ?", session.Username).Delete(&model.Session{})
	database.Save(session)
	if database.Error != nil {
		return database.Error
	}
	return nil
}

func EndSession(sessionID string) error {
	database.Where("id = ?", sessionID).Delete(&model.Session{})
	if database.Error != nil {
		return database.Error
	}
	return nil
}

func GetSession(sessionID string) (model.Session, error) {

	var session model.Session
	database.Where("id = ?", sessionID).Limit(1).Find(&session)

	if session.ID == "" {
		return model.Session{}, errors.New("invalid session")
	}

	if session.Expires.Sub(time.Now()) < 0 {
		database.Delete(session)
		return model.Session{}, nil
	}

	return session, nil
}
