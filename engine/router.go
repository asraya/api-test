package engine

import (
	"net/http"

	"api-test/app/handler"
	"api-test/app/repository"
	"api-test/config"
	"api-test/helpers"
	"api-test/service"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupConnection()

	newsRepository repository.NewsRepository = repository.NewNewsRepository(db)
	newsService    service.NewsService       = service.NewNewsService(newsRepository)
	newsHandler    handler.NewsHandler       = handler.NewNewsHandler(newsService)

	tagRepository repository.TagRepository = repository.NewTagRepository(db)
	tagService    service.TagService       = service.NewTagService(tagRepository)
	tagHandler    handler.TagHandler       = handler.NewTagHandler(tagService)
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())

	// CORS
	r.Use(CORSMiddleware())

	// Routes
	v1 := r.Group("api/v1")
	{
		routes := v1.Group("/")
		{
			n := routes.Group("/news")
			{
				n.GET("", func(context *gin.Context) {
					code := http.StatusOK
					pagination := helpers.GeneratePaginationRequest(context)
					response := newsService.PaginationNews(newsRepository, context, pagination)
					if !response.Status {
						code = http.StatusBadRequest
					}
					context.JSON(code, response)
				})
				n.POST("/", newsHandler.Insert)
				n.GET("/:id", newsHandler.FindByID)
				n.PUT("/:id", newsHandler.Update)
				n.DELETE("/:id", newsHandler.Delete)
			}
			t := routes.Group("/tags")
			{
				t.GET("", func(context *gin.Context) {
					code := http.StatusOK
					pagination := helpers.GeneratePaginationRequest(context)
					response := tagService.PaginationTag(tagRepository, context, pagination)
					if !response.Status {
						code = http.StatusBadRequest
					}
					context.JSON(code, response)
				})
				t.POST("/", tagHandler.Insert)
				t.GET("/:id", tagHandler.FindByID)
				t.PUT("/:id", tagHandler.Update)
				t.DELETE("/:id", tagHandler.Delete)
			}
		}
	}
	return r
}

// CORSMiddleware ..
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
