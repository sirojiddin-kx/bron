package api

import (
	"github.com/sirojiddin-kx/bron/config"
	"github.com/sirojiddin-kx/bron/pkg/logger"
	"github.com/sirojiddin-kx/bron/storage"

	v1 "github.com/sirojiddin-kx/bron/api/handlers/v1"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/sirojiddin-kx/bron/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouterOptions struct {
	Log     logger.Logger
	Cfg     *config.Config
	Storage storage.StorageI
}

func New(opt *RouterOptions) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	cfg := cors.DefaultConfig()

	cfg.AllowHeaders = append(cfg.AllowHeaders, "*")
	cfg.AllowAllOrigins = true
	cfg.AllowCredentials = true

	router.Use(cors.New(cfg))

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Log:     opt.Log,
		Cfg:     opt.Cfg,
		Storage: opt.Storage,
	})

	apiV1 := router.Group("/v1")

	{
		// Company Service
		apiV1.POST("/company-service", handlerV1.CompanyServcieCreate)
		apiV1.GET("/company-service", handlerV1.ListCompanyService)

		// Employee
		apiV1.POST("/employee", handlerV1.EmployeeCreate)
		apiV1.GET("/employee", handlerV1.ListEmployee)
		apiV1.GET("/employee/:employee_id", handlerV1.GetEmployeeByID)

		// Client 
		apiV1.POST("/client", handlerV1.ClientCreate)
		apiV1.GET("/clients", handlerV1.ListClients)
	}

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
