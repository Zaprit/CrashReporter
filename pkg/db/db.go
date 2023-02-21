package db

import (
	"errors"
	"time"

	"github.com/Zaprit/CrashReporter/pkg/model"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB

func OpenDB(path string) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic(err.Error())
	}
	database = db
}

func MigrateDB() {
	err := database.AutoMigrate(
		&model.Notice{},
		&model.Report{},
		&model.ReportCategory{},
		&model.ReportType{},
		&model.Session{},
		&model.Comment{},
	)
	if err != nil {
		panic(err.Error())
	}
}

func GetNotifications() []model.Notice {
	var Notices []model.Notice
	database.Where("ended = 0").Find(&Notices)
	return Notices
}

func GetReport(uuid string) model.Report {
	var report model.Report
	database.Where("uuid = ?", uuid).First(&report)
	return report
}

func GetReportID(uuid string) (uint, error) {
	var report model.Report
	database.Where("uuid = ?", uuid).First(&report)
	if database.RowsAffected == 0 {
		return 0, errors.New("report not found")
	}
	return report.ID, nil
}

func ReadReport(uuid string) {
	database.Model(&model.Report{}).Where("uuid = ?", uuid).Update("read", true)
}

func UnreadReport(uuid string) {
	database.Model(&model.Report{}).Where("uuid = ?", uuid).Update("read", false)
}

func GetReports() []model.Report {
	var reports []model.Report
	database.Find(&reports)
	return reports
}

func GetReportCategories() map[string][]string {
	var categories []model.ReportCategory
	database.Model(&model.ReportCategory{}).Preload("Types").Find(&categories)

	var defaultCategories = model.DefaultCategories

	for _, category := range categories {
		if category.Archived {
			continue
		}

		var types = make([]string, len(category.Types))
		for key, reportType := range category.Types {
			if reportType.Archived {
				continue
			}
			types[key] = reportType.Name
		}
		defaultCategories[category.Name] = types
	}

	return defaultCategories
}

func GetComments(reportID uint) []model.Comment {
	var comments []model.Comment
	database.Where("report_id = ?", reportID).Find(&comments)
	return comments
}

func PostComment(comment model.Comment) {
	database.Save(comment)
}

func ReportTypeExists(reportType string) bool {
	categories := GetReportCategories()

	var exists = false

	for _, category := range categories {
		for _, t := range category {
			if t == reportType {
				exists = true
			}
		}
	}

	return exists
}

func SubmitReport(report *model.Report) error {
	database.Save(report)
	if database.Error != nil {
		return database.Error
	}
	return nil
}

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

//func Login(username string, password string) model.User
