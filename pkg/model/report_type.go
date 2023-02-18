package model

type ReportType struct {
	ID   uint `gorm:"primarykey"`
	Name string
    Archived bool
    ReportCategoryID uint
}
