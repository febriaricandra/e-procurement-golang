package models

type ItemsVendor struct {
	BaseModel
	UsersItemID uint
	VendorID    uint
	UsersItem   UsersItem `gorm:"foreignKey:UsersItemID;references:ID"`
	Vendor      Vendor    `gorm:"foreignKey:VendorID;references:ID"`
}
