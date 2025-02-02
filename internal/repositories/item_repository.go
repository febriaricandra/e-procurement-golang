package repositories

import (
	"my-procurement-system/internal/models"

	"gorm.io/gorm"
)

type ItemRepository interface {
	CreateItem(item *models.Item) error
	GetItemByID(id uint) (*models.Item, error)
	GetAllItems() ([]models.Item, error)
}

type itemRepo struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *itemRepo {
	return &itemRepo{
		db: db,
	}
}

func (i *itemRepo) CreateItem(item *models.Item) error {
	return i.db.Create(item).Error
}

func (i *itemRepo) GetItemByID(itemID uint) (*models.Item, error) {
	var item models.Item
	err := i.db.First(&item, itemID).Error
	return &item, err
}

func (i *itemRepo) GetAllItems() ([]models.Item, error) {
	var items []models.Item
	err := i.db.Find(&items).Error

	return items, err
}
