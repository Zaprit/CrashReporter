package importer

// OldReport contains all the fields from the original CrashHelperCore
type OldReport struct {
	Title         string `json:"title"`
	Username      string `json:"username"`
	IssueType     string `json:"issueType"`
	IssuePriority string `json:"issuePriority"`
	Platform      string `json:"platform"`
	Details       string `json:"details"`
	HasEvidence   bool   `json:"hasEvidence"`
	IpAddress     string `json:"ipAddress"`
	Timestamp     string `json:"timestamp"`
}
