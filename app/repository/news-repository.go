package repository

import (
	"api-test/dto"
	"api-test/models"
	"fmt"
	"math"
	"strings"

	"gorm.io/gorm"
)

//NewsRepository is a ....
type NewsRepository interface {
	InsertNews(m models.News) models.News
	UpdateNews(m models.News) models.News
	DeleteNews(m models.News)
	AllNews() []models.News
	FindNewsByID(newsID uint64) models.News
	PaginationNews(pagination *dto.Pagination) (RepositoryResult, int)
}

type newsConnection struct {
	connection *gorm.DB
}

//NewNewsRepository
func NewNewsRepository(dbConn *gorm.DB) NewsRepository {
	return &newsConnection{
		connection: dbConn,
	}
}

func (db *newsConnection) InsertNews(m models.News) models.News {
	db.connection.Save(&m)
	db.connection.Preload("Tag").Find(&m)
	return m
}

func (db *newsConnection) UpdateNews(m models.News) models.News {
	db.connection.Save(&m)
	db.connection.Preload("Topic").Find(&m)
	return m
}

func (db *newsConnection) DeleteNews(m models.News) {
	db.connection.Delete(&m)
}

func (db *newsConnection) FindNewsByID(newsID uint64) models.News {
	var news models.News
	db.connection.Preload("Topic").Find(&news, newsID)
	return news
}

func (db *newsConnection) AllNews() []models.News {
	var newss []models.News
	db.connection.Preload("Topic").Find(&newss)
	return newss
}
func (db *newsConnection) PaginationNews(pagination *dto.Pagination) (RepositoryResult, int) {

	var newsy []models.News
	var count int64

	totalNews, totalRows, totalPages, fromRow, toRow, toPro := 0, 0, 0, 0, 0, 0

	offset := pagination.Page * pagination.Limit

	// get data with limit, offset & order
	find := db.connection.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

	// generate where query
	searchs := pagination.Searchs

	if searchs != nil {
		for _, value := range searchs {
			column := value.Column
			action := value.Action
			query := value.Query

			switch action {
			case "equals":
				whereQuery := fmt.Sprintf("%s = ?", column)
				find = find.Where(whereQuery, query)
				break
			case "contains":
				whereQuery := fmt.Sprintf("%s LIKE ?", column)
				find = find.Where(whereQuery, "%"+query+"%")
				break
			case "in":
				whereQuery := fmt.Sprintf("%s IN (?)", column)
				queryArray := strings.Split(query, ",")
				find = find.Where(whereQuery, queryArray)
				break

			}
		}
	}

	find = find.Find(&newsy)
	// has error find data
	errFind := find.Error

	if errFind != nil {
		return RepositoryResult{Error: errFind}, totalPages
	}

	pagination.Rows = newsy
	// count all data
	errCount := db.connection.Model(&models.News{}).Count(&count).Error

	if errCount != nil {
		return RepositoryResult{Error: errCount}, totalPages
	}

	pagination.TotalRows = int(count)

	// calculate total pages
	totalPages = int(math.Ceil(float64(count)/float64(pagination.Limit))) - 1
	if pagination.Page == 0 {
		// set from & to row on first page
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			// calculate from & to row
			fromRow = pagination.Page*pagination.Limit + 1
			toRow = (pagination.Page + 1) * pagination.Limit
		}
	}

	if toRow > int(count) {
		// set to row with total rows
		toRow = totalRows

	}

	// count all News
	errCountNews := db.connection.Model(&models.News{}).Count(&count).Error

	if errCountNews != nil {
		return RepositoryResult{Error: errCountNews}, totalNews
	}

	// calculate total pages
	totalNews = int(math.Ceil(float64(count)/float64(pagination.Limit))) - 1
	if pagination.Page == 0 {
		// set from & to row on first page
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalNews {
			// calculate from & to row
			fromRow = pagination.Page*pagination.Limit + 1
			toRow = (pagination.Page + 1) * pagination.Limit
		}
	}

	if toPro > int(count) {
		// set to row with total rows
		toPro = totalNews
	}

	pagination.FromRow = fromRow
	pagination.ToRow = toRow

	return RepositoryResult{Result: pagination}, totalPages
}
