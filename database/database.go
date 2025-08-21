package database

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
    database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    DB = database
}
