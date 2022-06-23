package service

import (
	"api-test/app/repository"
	"api-test/dto"
	"api-test/models"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

//TagService is a ....
type TagService interface {
	Insert(b dto.TagCreateDTO) models.Tag
	Update(b dto.TagUpdateDTO) models.Tag
	Delete(b models.Tag)
	All() []models.Tag
	FindByID(tagID uint64) models.Tag
	IsAllowedToEdit(tagID uint64) bool
	PaginationTag(repo repository.TagRepository, context *gin.Context, pagination *dto.Pagination) dto.Response
}

type tagService struct {
	tagRepository repository.TagRepository
}

//NewTagService .....
func NewTagService(tagRepo repository.TagRepository) TagService {
	return &tagService{
		tagRepository: tagRepo,
	}
}

func (service *tagService) Insert(b dto.TagCreateDTO) models.Tag {
	tag := models.Tag{}
	err := smapping.FillStruct(&tag, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.tagRepository.InsertTag(tag)
	return res
}
func (service *tagService) IsAllowedToEdit(tagID uint64) bool {
	b := service.tagRepository.FindTagByID(tagID)
	id := (b.ID)
	return tagID == id
}

func (service *tagService) Update(b dto.TagUpdateDTO) models.Tag {
	tag := models.Tag{}
	err := smapping.FillStruct(&tag, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.tagRepository.UpdateTag(tag)
	return res
}

func (service *tagService) Delete(b models.Tag) {
	service.tagRepository.DeleteTag(b)
}

func (service *tagService) All() []models.Tag {
	return service.tagRepository.AllTag()
}

func (service *tagService) FindByID(tagID uint64) models.Tag {
	return service.tagRepository.FindTagByID(tagID)
}

func (service *tagService) PaginationTag(repo repository.TagRepository, context *gin.Context, pagination *dto.Pagination) dto.Response {

	operationResult, totalPages := repo.PaginationTag(pagination)

	if operationResult.Error != nil {
		return dto.Response{Status: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*dto.Pagination)

	urlPath := context.Request.URL.Path

	searchQueryParams := ""

	for _, search := range pagination.Searchs {
		searchQueryParams += fmt.Sprintf("&%s.%s=%s", search.Column, search.Action, search.Query)
	}

	data.FirstPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, 0, pagination.Sort) + searchQueryParams
	data.LastPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, totalPages, pagination.Sort) + searchQueryParams

	if data.Page > 0 {
		data.PreviousPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, data.Page-1, pagination.Sort) + searchQueryParams
	}

	if data.Page < totalPages {
		data.NextPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, data.Page+1, pagination.Sort) + searchQueryParams
	}

	if data.Page > totalPages {
		data.PreviousPage = ""
	}

	return dto.Response{Status: true, Data: data}
}
