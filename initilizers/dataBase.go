package initializers

import (
	"e-com-be-go/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/e-com-go?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("unable to connect with DB")
	}

	// Migrate the schema
	err = DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Product{}, &models.Cart{}, &models.CartProduct{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migrated successfully")
}
