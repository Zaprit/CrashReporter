// Package importer is used to convert old report json files from CrashHelperCore into reports to go into a database
package importer

import (
	"encoding/json"
	"github.com/Zaprit/CrashReporter/pkg/db"
	"github.com/Zaprit/CrashReporter/pkg/lighthouse_client"
	"github.com/Zaprit/CrashReporter/pkg/model"
	"github.com/Zaprit/CrashReporter/pkg/model/importer"
	"log"
	"net/http"
	"os"
	"strings"
)

func importReport(jsonReport []byte) (importer.OldReport, error) {
	var report importer.OldReport
	err := json.Unmarshal(jsonReport, &report)
	return report, err
}

func convertReport(report importer.OldReport, reportName string) model.Report {
	lhUser, status, err := lighthouse_client.GetUser(report.Username)
	if err != nil || status != http.StatusOK {
		log.Printf("Failed to get lighthouse user for report: %s", reportName)
	}

	timestamp, err := parseOrdinalDate("January 2 2006 3:4:35 pm", report.Timestamp)
	if err != nil {
		log.Printf("Failed to parse date on report: %s", reportName)
	}

	newReport := model.Report{
		UUID:             reportName,
		Title:            "(Imported) " + report.Title,
		Username:         report.Username,
		UserID:           lhUser.UserId,
		Type:             report.IssueType,
		Platform:         report.Platform,
		Description:      report.Details,
		Evidence:         report.HasEvidence,
		SubmitterIP:      report.IpAddress,
		SubmitTime:       timestamp,
		Comments:         nil,
		Priority:         report.IssuePriority,
		DiscordMessageID: "",
	}

	return newReport
}

// ImportOldReports takes all the reports from an installation of CrashHelperCore and submits them to the databases
func ImportOldReports() {
	dbFiles, err := os.ReadDir("./db/reports")
	if err != nil {
		return
	}

	for _, file := range dbFiles {
		if file.IsDir() {
			continue
		}
		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		reportName := strings.TrimSuffix(file.Name(), ".json")

		reportData, err := os.ReadFile("./db/reports/" + file.Name())
		if err != nil {
			log.Println("failed to import report: " + reportName)
			continue
		}

		report, err := importReport(reportData)
		if err != nil {
			log.Println("failed to deserialize report: " + reportName)
			continue
		}

		convertedReport := convertReport(report, reportName)

		err = db.SaveReport(&convertedReport)
		if err != nil {
			log.Printf("failed to save imported report %s(%s)", report.Title, convertedReport.UUID)
		}
	}
}
