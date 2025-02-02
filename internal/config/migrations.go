package config

import (
	"log"
	"my-procurement-system/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Migrate(db *gorm.DB) error {
	log.Println("🔄 Memulai migrasi database...")

	db.Logger = db.Logger.LogMode(logger.Info)
	err := db.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Category{},
		&models.Warehouse{},
		&models.Item{},
		&models.Vendor{},
		&models.UsersItem{},
		&models.Approval{},
		&models.ItemsVendor{},
		// &models.AuditLog{}, // Jika sudah dibuat
	)

	if err != nil {
		log.Fatalf("❌ Gagal melakukan migrasi: %v", err)
		return err
	}

	log.Println("✅ Migrasi database berhasil!")
	return nil
}
