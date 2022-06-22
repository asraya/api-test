package handler

import (
	"api-test/dto"
	"api-test/helpers"
	"api-test/models"
	"api-test/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//TagHandler is a ...
type TagHandler interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type tagHandler struct {
	tagService service.TagService
}

func NewTagHandler(tagServ service.TagService) TagHandler {
	return &tagHandler{
		tagService: tagServ,
	}
}

func (c *tagHandler) All(context *gin.Context) {
	var tags []models.Tag = c.tagService.All()
	res := helpers.BuildResponse(true, "OK", tags)
	context.JSON(http.StatusOK, res)
}

func (c *tagHandler) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helpers.BuildErrorResponse("No param id was found", err.Error(), helpers.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var tag models.Tag = c.tagService.FindByID(id)
	if (tag == models.Tag{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helpers.BuildResponse(true, "OK", tag)
		context.JSON(http.StatusOK, res)
	}
}

func (c *tagHandler) Insert(context *gin.Context) {
	var tagCreateDTO dto.TagCreateDTO
	errDTO := context.ShouldBind(&tagCreateDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		result := c.tagService.Insert(tagCreateDTO)
		response := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *tagHandler) Update(context *gin.Context) {
	var projectUpdateDTO dto.TagUpdateDTO
	errDTO := context.ShouldBind(&projectUpdateDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	} else {
		result := c.tagService.Update(projectUpdateDTO)
		response := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *tagHandler) Delete(ctx *gin.Context) {
	var project models.Tag

	project_id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed tou get id", "No param id were found", helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	}
	project.ID = project_id

	userID := fmt.Sprintf("news_id")
	if c.tagService.IsAllowedToEdit(userID, project.ID) {

		c.tagService.Delete(project)
		res := helpers.BuildResponse(true, "Deleted", helpers.EmptyObj{})
		ctx.JSON(http.StatusOK, res)

	} else {
		response := helpers.BuildErrorResponse("You dont have permission", "You are not the owner", helpers.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)

	}

}
