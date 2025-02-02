package controllers

import (
	"my-procurement-system/internal/models"
	"my-procurement-system/internal/services"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type ItemController struct {
	itemService *services.ItemService
}

func NewItemController(itemService *services.ItemService) *ItemController {
	return &ItemController{
		itemService: itemService,
	}
}

func (ic *ItemController) CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ic.itemService.CreateItem(&item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"success": true})
}

func (ic *ItemController) GetAllItems(c *gin.Context) {
	items, err := ic.itemService.GetAllItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}

func (ic *ItemController) GetItemByID(c *gin.Context) {
	itemID := c.Param("id")
	id, err := strconv.ParseUint(itemID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}
	item, err := ic.itemService.GetItemByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}
