package migration

import (
	"log"

	"api-test/database"
	"api-test/models"
)

func RunMigration() {
	db := database.PostsqlConn()

	if db.Error != nil {
		log.Fatalln(db.Error.Error())
	}

	db.AutoMigrate(&models.Tag{})
	db.AutoMigrate(&models.News{})
	db.AutoMigrate(&models.Topic{})

}
