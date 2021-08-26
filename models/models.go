package models

import (
	"github.com/ReeceDonovan/uni-bot/database"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitModels migrates modesls and initiates database connection
func InitModels() {
	db = database.InitDB()

	// db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";") // enable uuid generation on server
	db.AutoMigrate(&User{}, &Server{})
}

// User structure containing user level canvas token
type User struct {
	UID         string `gorm:"primaryKey"`
	CanvasToken string
	CreatedAt   int64 `gorm:"autoCreateTime"`
	UpdatedAt   int64 `gorm:"autoUpdateTime"`
}

// Server structure containing server level canvas token
type Server struct {
	SID         string `gorm:"primaryKey"`
	CanvasToken string
	CreatedAt   int64 `gorm:"autoCreateTime"`
	UpdatedAt   int64 `gorm:"autoUpdateTime"`
}

// GenericResponse contains a generic string
type GenericResponse struct {
	Data string `json:"data,omitempty"`
}
