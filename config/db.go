package config

import (
        "os"
        "restapi/models"

        "gorm.io/driver/postgres"
        "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
        db, err := gorm.Open(postgres.New(postgres.Config{
                DSN: os.Getenv("DATABASE_URL"),
        }), &gorm.Config{})
        if err != nil {
                panic(err)
        }

        db.AutoMigrate(&models.Category{})
        db.AutoMigrate(&models.Book{})
        db.AutoMigrate(&models.User{})
        DB = db
}

