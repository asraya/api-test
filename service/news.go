package service

import (
	"fmt"
	"log"

	"api-test/app/repository"
	"api-test/dto"
	"api-test/models"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

//NewsService is a
type NewsService interface {
	Insert(b dto.NewsCreateDTO) models.News
	Update(b dto.NewsUpdateDTO) models.News
	Delete(b models.News)
	All() []models.News
	FindByID(newsID uint64) models.News
	IsAllowedToEdit(userID string, newsID uint64) bool
	PaginationNews(repo repository.NewsRepository, context *gin.Context, pagination *dto.Pagination) dto.Response
}

type newsService struct {
	newsRepository repository.NewsRepository
}

//NewNewsService
func NewNewsService(newsRepo repository.NewsRepository) NewsService {
	return &newsService{
		newsRepository: newsRepo,
	}
}

func (service *newsService) Insert(b dto.NewsCreateDTO) models.News {
	news := models.News{}
	err := smapping.FillStruct(&news, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.newsRepository.InsertNews(news)
	return res
}

func (service *newsService) Update(b dto.NewsUpdateDTO) models.News {
	news := models.News{}
	err := smapping.FillStruct(&news, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.newsRepository.UpdateNews(news)
	return res
}

func (service *newsService) Delete(b models.News) {
	service.newsRepository.DeleteNews(b)
}

func (service *newsService) All() []models.News {
	return service.newsRepository.AllNews()
}

func (service *newsService) FindByID(newsID uint64) models.News {
	return service.newsRepository.FindNewsByID(newsID)
}

func (service *newsService) IsAllowedToEdit(tagsID string, newsID uint64) bool {
	b := service.newsRepository.FindNewsByID(newsID)
	id := fmt.Sprintf("%v", b.CreatedBy)
	return tagsID == id
}
func (service *newsService) PaginationNews(repo repository.NewsRepository, context *gin.Context, pagination *dto.Pagination) dto.Response {

	operationResult, totalPages := repo.PaginationNews(pagination)

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
