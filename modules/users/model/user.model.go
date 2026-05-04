package model

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"type:varchar(255)"`
	Name     string `gorm:"type:varchar(255)"`
	Password string `gorm:"type:varchar(255)"`
}
