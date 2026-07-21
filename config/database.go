package config

import (
	"log"
	"os"

	"github.com/okakafavour/supermarket-pos-backend/internal/category"
	"github.com/okakafavour/supermarket-pos-backend/internal/customer"
	"github.com/okakafavour/supermarket-pos-backend/internal/inventory"
	"github.com/okakafavour/supermarket-pos-backend/internal/payment"
	"github.com/okakafavour/supermarket-pos-backend/internal/product"
	"github.com/okakafavour/supermarket-pos-backend/internal/sale"
	"github.com/okakafavour/supermarket-pos-backend/internal/supplier"
	"github.com/okakafavour/supermarket-pos-backend/internal/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully")

	err = db.AutoMigrate(
		&user.User{},
		&category.Category{},
		&supplier.Supplier{},
		&product.Product{},
		&customer.Customer{},
		&sale.Sale{},
		&sale.SaleItem{},
		&inventory.InventoryLog{},
		&payment.Payment{},
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Database migrated successfully")

	return db
}
