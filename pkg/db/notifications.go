package db

import "github.com/Zaprit/CrashReporter/pkg/model"

func GetNotifications() []model.Notice {
	var Notices []model.Notice
	database.Where("ended = ?", false).Find(&Notices)
	return Notices
}
