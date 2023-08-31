package db

import (
	"fmt"
	"github.com/Zaprit/CrashReporter/pkg/model"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/netip"
	"time"
)

var database *gorm.DB

const TypeSqlite = "sqlite"
const TypeMysql = "mysql"
const TypePostgres = "postgres"

func OpenSQLiteDB(path string) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic(err.Error())
	}
	database = db
}

func OpenMySQLDB(host string, port uint16, databaseName string, user string, pass string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, databaseName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	database = db
}

func OpenPostgreSQLDB(host string, port uint16, databaseName string, user string, pass string) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		host, user, pass, databaseName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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

func BanIP(addr netip.Addr, reason string, until time.Time) {
	database.Save(&model.Ban{
		IP:      addr,
		Reason:  reason,
		EndDate: until,
	})
}

func IsBanned(addr netip.Addr) bool {
	ban := model.Ban{}
	database.Where("ip = ?", addr).First(&ban)
	if time.Now().Before(ban.EndDate) {
		return true
	}
	return false
}
