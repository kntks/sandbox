package sql

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Customer struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Main() {
	user := "docker"
	pass := "docker"
	host := "db"
	dbname := "test_database"
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, dbname)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Create
	db.Create(&Customer{Name: "test name"})
}
