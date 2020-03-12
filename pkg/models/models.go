package models

// User represent user table
type User struct {
	Name     string `gorm:"size:100;not null"`
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string `gorm:"size:255;not null"`
}
