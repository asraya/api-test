package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var (
	psqlConn *gorm.DB
	err      error
)

// initialize database
func init() {
	setupPostsqlConn()
}

func setupPostsqlConn() {
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	dns := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASS"))
	psqlConn, err = gorm.Open("postgres", dns)

	err = psqlConn.DB().Ping()
	if err != nil {
		panic(err)
	}

	psqlConn.LogMode(true)
}

// PostsqlConn return mysql connection from gorm ORM
func PostsqlConn() *gorm.DB {
	return psqlConn
}
