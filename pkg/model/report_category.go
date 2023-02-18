package model

var DefaultCategories = map[string][]string{
    "General Beacon Reports": {
        "Connection issues, error code while connecting, etc.",
        "Certain features are not working, i.e. cannot comment",
        "Gameplay issues, i.e. cannot dive in, broken peer to peer",
        "Website issues, cannot use certain features, or offline",
    },
    "Other Report Types": {
        "LBP issue, i.e. game crash, graphics bugs, etc.",
    },
}

type ReportCategory struct {
    ID int `gorm:"primarykey"`
    Name string
    Archived bool
    Types []ReportType
}