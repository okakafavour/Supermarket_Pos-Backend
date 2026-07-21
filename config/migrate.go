package config

import (
	"log"

	"github.com/okakafavour/supermarket-pos-backend/internal/category"
	"github.com/okakafavour/supermarket-pos-backend/internal/customer"
	"github.com/okakafavour/supermarket-pos-backend/internal/inventory"
	"github.com/okakafavour/supermarket-pos-backend/internal/payment"
	"github.com/okakafavour/supermarket-pos-backend/internal/product"
	"github.com/okakafavour/supermarket-pos-backend/internal/sale"
	"github.com/okakafavour/supermarket-pos-backend/internal/supplier"
	"github.com/okakafavour/supermarket-pos-backend/internal/user"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
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
}
