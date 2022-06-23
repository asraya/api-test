package handler

import (
	"net/http"
	"strconv"

	"api-test/dto"
	"api-test/helpers"
	"api-test/models"
	"api-test/service"

	"github.com/gin-gonic/gin"
)

//NewsHandler is a ...
type NewsHandler interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type newsHandler struct {
	newsService service.NewsService
}

func NewNewsHandler(newsServ service.NewsService) NewsHandler {
	return &newsHandler{
		newsService: newsServ,
	}
}

func (c *newsHandler) All(context *gin.Context) {
	var newss []models.News = c.newsService.All()
	res := helpers.BuildResponse(true, "OK", newss)
	context.JSON(http.StatusOK, res)
}

func (c *newsHandler) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helpers.BuildErrorResponse("No param id was found", err.Error(), helpers.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var news models.News = c.newsService.FindByID(id)
	if &news != &news {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helpers.BuildResponse(true, "OK", news)
		context.JSON(http.StatusOK, res)
	}
}

func (c *newsHandler) Insert(context *gin.Context) {
	var newsCreateDTO dto.NewsCreateDTO
	errDTO := context.ShouldBind(&newsCreateDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		result := c.newsService.Insert(newsCreateDTO)
		response := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *newsHandler) Update(context *gin.Context) {
	var newsUpdateDTO dto.NewsUpdateDTO

	errDTO := context.ShouldBind(&newsUpdateDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	if c.newsService.IsAllowedToEdit(newsUpdateDTO.ID) {
		result := c.newsService.Update(newsUpdateDTO)
		response := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *newsHandler) Delete(context *gin.Context) {
	var news models.News
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed tou get id", "No param id were found", helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	news.ID = id
	if c.newsService.IsAllowedToEdit(news.ID) {
		c.newsService.Delete(news)
		res := helpers.BuildResponse(true, "Deleted", helpers.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helpers.BuildErrorResponse("You dont have permission", "You are not the owner", helpers.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}
