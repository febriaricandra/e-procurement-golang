package routes

import (
	"my-procurement-system/internal/controllers"
	"my-procurement-system/internal/repositories"
	"my-procurement-system/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterItemRoutes(r *gin.RouterGroup, db *gorm.DB) {
	itemRepo := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepo)
	itemController := controllers.NewItemController(itemService)

	itemGroup := r.Group("/items")
	{
		itemGroup.POST("/", itemController.CreateItem)
		itemGroup.GET("/:id", itemController.GetItemByID)
		itemGroup.GET("/", itemController.GetAllItems)
	}
}
