package config

import (
	"article-service/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbUser := getEnv("DB_USER", "root")
	dbPass := getEnv("DB_PASSWORD", "")
	dbHost := getEnv("DB_HOST", "127.0.0.1")
	dbPort := getEnv("DB_PORT", "3306")
	dbName := getEnv("DB_NAME", "db_article")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = migratePosts(database)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = database
	fmt.Println("Database connected and migrated successfully!")
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// migratePosts preserves data from the legacy titles column while allowing
// inserts that use the current title field. AutoMigrate retains old columns.
func migratePosts(db *gorm.DB) error {
	migrator := db.Migrator()
	if migrator.HasTable(&models.Post{}) && migrator.HasColumn(&models.Post{}, "titles") {
		if !migrator.HasColumn(&models.Post{}, "title") {
			if err := migrator.RenameColumn(&models.Post{}, "titles", "title"); err != nil {
				return err
			}
		} else {
			legacy := struct {
				Titles *string `gorm:"column:titles;type:varchar(200)"`
			}{}
			if err := db.Table("posts").Migrator().AlterColumn(&legacy, "Titles"); err != nil {
				return err
			}
		}
	}
	return db.AutoMigrate(&models.Post{})
}
