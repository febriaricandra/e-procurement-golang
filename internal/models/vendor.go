package models

type Vendor struct {
	BaseModel
	Name string `gorm:"type:varchar(255)"`
}
