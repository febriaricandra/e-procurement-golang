package models

type Warehouse struct {
	BaseModel
	Name string `gorm:"type:varchar(255)"`
}
