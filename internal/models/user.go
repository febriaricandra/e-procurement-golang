package models

type User struct {
	BaseModel
	Name         string `gorm:"type:varchar(255);unique;not null"`
	Email        string `gorm:"type:varchar(255);unique;not null"`
	Password     string `gorm:"-"`
	RoleID       uint   `gorm:"not null"` //foreign key
	Role         Role   `gorm:"foreignKey:RoleID;references:ID"`
	RefreshToken []RefreshToken
}

type Role struct {
	BaseModel
	Name string `gorm:"type:varchar(255);not null"`
}
