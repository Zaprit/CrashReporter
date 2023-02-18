package model

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Username string
    Password string
}

type UserFormData struct {
    Username string
    PasswordHash string
}