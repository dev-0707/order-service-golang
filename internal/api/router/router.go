package router

import (
	// "io"
	"io"
	"order-service/internal/api/middlewares"
	"order-service/internal/api/users"
	"order-service/internal/pkg/config"
	"order-service/internal/pkg/persistence"
	"order-service/internal/pkg/service"
	http_err "order-service/pkg/http-err"
	"os"

	swaggerFiles "github.com/swaggo/"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/gin-gonic/gin"
)

//https://github.com/antonioalfa22/go-rest-template/blob/master/internal/api/router/router.go
func Setup(config *config.Configuration) *gin.Engine {
	ginEngine := createGinEngine(config)
	userController := buildUserController()
	// Routes
	addSwaggerDocs(ginEngine, config)
	addMiddlewares(ginEngine, config)
	addUserController(ginEngine, config, userController)
	if config.Server.LogFileEnabled {
		addLogToFile(config)
	}
	return ginEngine
}

func addUserController(ginEngine *gin.Engine, config *config.Configuration, userController *users.UsersControllerDelegate) {
	// v1 := router.Group("/api/v1")
	// +       {
	// +               v1.GET("/books", handlers.GetBooks)
	// +               v1.GET("/books/:isbn", handlers.GetBookByISBN)
	// +               // router.DELETE("/books/:isbn", handlers.DeleteBookByISBN)
	// +               // router.PUT("/books/:isbn", handlers.UpdateBookByISBN)
	// +               v1.POST("/books", handlers.PostBook)
	// +       }

	users.RegisterHandlers(ginEngine, userController)
}

func createGinEngine(config *config.Configuration) *gin.Engine {
	gin.SetMode(config.Server.Mode)
	if gin.Mode() == "release" {
		gin.DefaultWriter = io.Discard
	}
	ginEngine := gin.New()
	return ginEngine
}

func addSwaggerDocs(r *gin.Engine, configs *config.Configuration) {
	apiConfig := configs.Server.Api
	r.StaticFile(apiConfig.DocsPath, apiConfig.DocsPath)

	ginSwaggerConfigs := &ginSwagger.Config{
		URL: apiConfig.DocsPath,
	}

	r.GET(apiConfig.DocsUrl, ginSwagger.CustomWrapHandler(ginSwaggerConfigs, swaggerFiles.Handler))
}

func buildUserController() *users.UsersControllerDelegate {
	userRepository := persistence.GetUserRepository()

	userService := service.NewUserService(*userRepository)
	userController := users.NewUsersControllerDelegate(*userService)
	return userController
}

func addMiddlewares(ginEngine *gin.Engine, configs *config.Configuration) {
	// app.Use(gin.Recovery())
	ginEngine.Use(middlewares.CORS())
	ginEngine.NoRoute(middlewares.NoRouteHandler())
	ginEngine.Use(gin.CustomRecovery(http_err.InternalServerErrorHandler))

	if configs.Server.LogEnabled {
		ginEngine.Use(middlewares.DefaultStructuredLogger())
	}
}

func addLogToFile(configs *config.Configuration) {
	f, _ := os.Create(configs.Server.LogPath)
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f)
}
