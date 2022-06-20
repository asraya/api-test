package config

import (
	"net/http"

	"api-test/app/handler"
	"api-test/config"
	"api-test/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupConnection()

	newsRepository repo.NewsRepository = repo.NewNewsRepository(db)
	newsService    service.NewsService = service.NewNewsService(newsRepository)
	newsHandler    handler.NewsHandler = handler.NewNewsHandler(newsService, jwtService)
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
			pakets := routes.Group("/news")
			{
				news.GET("", func(context *gin.Context) {
					code := http.StatusOK
					pagination := helper.GeneratePaginationRequest(context)
					response := s.PaginationNews(newsRepo, context, pagination)
					if !response.Status {
						code = http.StatusBadRequest
					}
					context.JSON(code, response)
				})
				pakets.POST("/", newsHandler.Create)
				pakets.GET("/:id", newsHandler.Show)
				pakets.PUT("/:id", newsHandler.Update)
				pakets.DELETE("/:id", newsHandler.Delete)
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
