package db
import (
    "github.com/Zaprit/CrashReporter/pkg/model"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var database *gorm.DB

func OpenDB(path string) {
    db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
    if err != nil {
        panic(err.Error())
    }
    database = db
}

func MigrateDB() {
    err := database.AutoMigrate(&model.Notice{}, &model.Report{}, &model.ReportCategory{}, &model.ReportType{})
    if err != nil {
        panic(err.Error())
    }
}

func GetNotifications() []model.Notice {
    var Notices []model.Notice
    database.Where("ended = 0").Find(&Notices)
    return Notices
}

func GetReport(id string) model.Report {
    var Report model.Report
    database.Where("id = ?", id).First(&Report)
    return Report
}



func GetReportCategories() map[string][]string {
    var categories []model.ReportCategory
    database.Model(&model.ReportCategory{}).Preload("Types").Find(&categories)

    var defaultCategories = model.DefaultCategories

    for _,category := range categories {
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

func ReportTypeExists(reportType string) bool{
    categories := GetReportCategories()

    var exists = false

    for _,category := range categories {
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

//func Login(username string, password string) model.User