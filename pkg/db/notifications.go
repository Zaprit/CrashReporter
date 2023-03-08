package db

import "github.com/Zaprit/CrashReporter/pkg/model"

func GetNotifications() []model.Notice {
	var Notices []model.Notice
	database.Where("ended = 0").Find(&Notices)
	return Notices
}
