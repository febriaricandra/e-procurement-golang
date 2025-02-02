package services

import (
	"my-procurement-system/internal/models"
	"my-procurement-system/internal/repositories"
)

type ItemService struct {
	itemRepo repositories.ItemRepository
}

func NewItemService(repo repositories.ItemRepository) *ItemService {
	return &ItemService{
		itemRepo: repo,
	}
}

func (i ItemService) CreateItem(item *models.Item) error {
	return i.itemRepo.CreateItem(item)
}

func (i ItemService) GetAllItems() ([]models.Item, error) {
	return i.itemRepo.GetAllItems()
}

func (i ItemService) GetItemByID(itemID uint) (*models.Item, error) {
	return i.itemRepo.GetItemByID(itemID)
}
