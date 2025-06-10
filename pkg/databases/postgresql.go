package database

import (
	"e-com/modules/entities"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydatabase"
)

func ConnentDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Println("Database migration completed!")

	AutoMigrate(db)

	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&entities.User{},
		&entities.Product{},
		&entities.Image{},
		&entities.Category{},
		&entities.Brand{},
		&entities.Order{},
		&entities.Cart{},

		&entities.OrderItem{},
		&entities.CartItem{},
	)
}
