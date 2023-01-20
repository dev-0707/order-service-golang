package router

import (
	"io"
	"order-service/internal/api/middlewares"
	"order-service/internal/api/users"
	"order-service/internal/pkg/config"
	"order-service/internal/pkg/persistence"
	"order-service/internal/pkg/service"
	http_err "order-service/pkg/http-err"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

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
	swagger, err := users.GetSwagger()
	if err != nil {
		log.Error().Err(err).Msg("Errore configurazione Swagger Docs")
		os.Exit(1)
	}

	swagger.Servers = nil

	r.StaticFile(configs.Server.Api.DocsPath, configs.Server.Api.DocsPath)
	config := &ginSwagger.Config{
		URL: configs.Server.Api.DocsUrl,
	}
	r.GET(configs.Server.Api.SwaggerDocsUrl, ginSwagger.CustomWrapHandler(config, swaggerFiles.Handler))
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
