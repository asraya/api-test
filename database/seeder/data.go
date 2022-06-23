package seeder

import (
	"fmt"
	"time"

	"api-test/database"
	"api-test/models"
)

func Topic() {
	db := database.PostsqlConn()

	now := time.Now()

	var types []models.Topic

	var type2 = models.Topic{
		ID:        2,
		Name:      "Investment",
		CreatedAt: now,
		UpdatedAt: now,
	}
	types = append(types, type2)

	var type1 = models.Topic{
		ID:        3,
		Name:      "How to start Investment",
		CreatedAt: now,
		UpdatedAt: now,
	}
	types = append(types, type1)

	var type3 = models.Topic{
		ID:        4,
		Name:      "Mutual fund is safe Investment",
		CreatedAt: now,
		UpdatedAt: now,
	}
	types = append(types, type3)

	for _, Topic := range types {
		if err := db.Where("name = ?", Topic.Name).Find(&Topic).Error; err != nil {
			db.Create(&Topic)
		}
		fmt.Printf("package type %s has been created\n", Topic.Name)
	}
}

func News() {
	db := database.PostsqlConn()

	now := time.Now()

	var types []models.News

	var type2 = models.News{
		ID:        1,
		Name:      "Request News",
		Status:    "draft",
		CreatedAt: now,
		UpdatedAt: now,
	}
	types = append(types, type2)

	for _, News := range types {
		if err := db.Where("name = ?", News.Name).Find(&News).Error; err != nil {
			db.Create(&News)
		}
		fmt.Printf("type2 %s has been created\n", News.Name)
	}
}

func Tag() {
	db := database.PostsqlConn()

	now := time.Now()

	var types []models.Tag

	var type2 = models.Tag{
		ID:        2,
		Name:      "Save Investment",
		NewsId:    1,
		CreatedAt: now,
		UpdatedAt: now,
	}
	types = append(types, type2)

	var type1 = models.Tag{
		ID:        3,
		Name:      "Investment",
		NewsId:    1,
		CreatedAt: now,
		UpdatedAt: now,
	}
	types = append(types, type1)

	var type3 = models.Tag{
		ID:        4,
		Name:      "Mutual Fund",
		NewsId:    1,
		CreatedAt: now,
		UpdatedAt: now,
	}
	types = append(types, type3)

	for _, Tag := range types {
		if err := db.Where("name = ?", Tag.Name).Find(&Tag).Error; err != nil {
			db.Create(&Tag)
		}
		fmt.Printf("package type %s has been created\n", Tag.Name)
	}
}

// Package ..

func RunSeeder() {
	Tag()
	News()
	Topic()

}
