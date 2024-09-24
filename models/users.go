package models

type Users struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(50)" json:"name" binding:"required,min=3"`
	Email    string `gorm:"type:text" json:"email" binding:"required,min=3,unique"`
	Password string `gorm:"type:text" json:"password" binding:"required,min=3"`
}
