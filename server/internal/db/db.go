package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewMySQL membuka koneksi GORM ke MySQL DSN
func NewMySQL(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("[db] gagal konek ke MySQL: %v", err) // fail-fast di startup
	}
	return db
}