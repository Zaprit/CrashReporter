package db

import "github.com/Zaprit/CrashReporter/pkg/model"

func GetReport(uuid string) model.Report {
	var report model.Report
	database.Where("uuid = ?", uuid).First(&report)
	return report
}

func GetReportID(uuid string) (uint, error) {
	var report model.Report
	database.Where("uuid = ?", uuid).First(&report)
	if database.Error != nil {
		return 0, database.Error
	}

	return report.ID, nil
}

func ReadReport(uuid string) {
	database.Model(&model.Report{}).Where("uuid = ?", uuid).Update("read", true)
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
	database.Where("report_id = ?", reportID).Order("create_time DESC").Find(&comments)
	return comments
}

func PostComment(comment model.Comment) {
	database.Save(&comment)
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

func SaveReport(report *model.Report) error {
	database.Save(report)
	if database.Error != nil {
		return database.Error
	}
	return nil
}
