package model

type LighthouseUser struct {
    UserId uint
    Username string
    EmailAddressVerified bool
    IconHash string
    Biography string
    YayHash string
    BooHash string
    MehHash string
    LastLogin int
    LastLogout int
    LevelVisibility int
    ProfileVisibility int
    CommentEnabled bool
}