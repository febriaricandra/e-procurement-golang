package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Item struct {
	BaseModel
	Name        string `gorm:"type:varchar(255)"`
	Price       float64
	Stock       int
	CategoryID  uint
	WarehouseID uint
	Category    Category  `gorm:"foreignKey:CategoryID;references:ID"`
	Warehouse   Warehouse `gorm:"foreignKey:WarehouseID;references:ID"`
}

func (i *Item) TableName() string {
	return "items"
}

type Approval struct {
	BaseModel
	ApprovedBy   uint
	UsersItemsID uint
	Status       string `gorm:"type:varchar(50);default:'pending'"`
	ApprovedAt   *time.Time
	User         User      `gorm:"foreignKey:ApprovedBy;references:ID"`
	UsersItem    UsersItem `gorm:"foreignKey:UsersItemsID;references:ID"`
}

func (a *Approval) TableName() string {
	return "approvals"
}

type UsersItem struct {
	BaseModel
	ItemID uint
	UserID uint
	Item   Item `gorm:"foreignKey:ItemID;references:ID"`
	User   User `gorm:"foreignKey:UserID;references:ID"`
}

func (u *UsersItem) TableName() string {
	return "users_items"
}
