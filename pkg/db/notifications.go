package db

import "github.com/Zaprit/CrashReporter/pkg/model"

func GetNotifications() []model.Notice {
	var Notices []model.Notice
	database.Where("ended = ?", false).Find(&Notices)
	return Notices
}

func PostNotice(notice model.Notice) {
	database.Save(&notice)
}

func DeleteNotice(noticeID string) {
	database.Delete(model.Notice{}, "id = ?", noticeID)
}
