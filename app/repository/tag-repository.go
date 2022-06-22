package repository

import (
	"api-test/dto"
	"api-test/models"
	"fmt"
	"math"
	"strings"

	"gorm.io/gorm"
)

//TagRepository is a ....
type TagRepository interface {
	InsertTag(m models.Tag) models.Tag
	UpdateTag(m models.Tag) models.Tag
	DeleteTag(m models.Tag)
	AllTag() []models.Tag
	FindTagByID(tagID uint64) models.Tag
	PaginationTag(pagination *dto.Pagination) (RepositoryResult, int)
}

type tagConnection struct {
	connection *gorm.DB
}

//NewTagRepository
func NewTagRepository(dbConn *gorm.DB) TagRepository {
	return &tagConnection{
		connection: dbConn,
	}
}

func (db *tagConnection) InsertTag(m models.Tag) models.Tag {
	db.connection.Save(&m)
	db.connection.Preload("News").Find(&m)
	return m
}

func (db *tagConnection) UpdateTag(m models.Tag) models.Tag {
	db.connection.Save(&m)
	db.connection.Preload("User").Find(&m)
	return m
}

func (db *tagConnection) DeleteTag(m models.Tag) {
	db.connection.Delete(&m)
}

func (db *tagConnection) FindTagByID(tagID uint64) models.Tag {
	var tag models.Tag
	db.connection.Preload("User").Find(&tag, tagID)
	return tag
}

func (db *tagConnection) AllTag() []models.Tag {
	var tags []models.Tag
	db.connection.Preload("User").Find(&tags)
	return tags
}
func (db *tagConnection) PaginationTag(pagination *dto.Pagination) (RepositoryResult, int) {

	var tagsy []models.Tag
	var count int64

	totalTag, totalRows, totalPages, fromRow, toRow, toPro := 0, 0, 0, 0, 0, 0

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

	find = find.Find(&tagsy)
	// has error find data
	errFind := find.Error

	if errFind != nil {
		return RepositoryResult{Error: errFind}, totalPages
	}

	pagination.Rows = tagsy
	// count all data
	errCount := db.connection.Model(&models.Tag{}).Count(&count).Error

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

	// count all Tag
	errCountTag := db.connection.Model(&models.Tag{}).Count(&count).Error

	if errCountTag != nil {
		return RepositoryResult{Error: errCountTag}, totalTag
	}
	pagination.TotalProject = int(count)

	// calculate total pages
	totalTag = int(math.Ceil(float64(count)/float64(pagination.Limit))) - 1
	if pagination.Page == 0 {
		// set from & to row on first page
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalTag {
			// calculate from & to row
			fromRow = pagination.Page*pagination.Limit + 1
			toRow = (pagination.Page + 1) * pagination.Limit
		}
	}

	if toPro > int(count) {
		// set to row with total rows
		toPro = totalTag
	}

	pagination.FromRow = fromRow
	pagination.ToRow = toRow

	return RepositoryResult{Result: pagination}, totalPages
}
