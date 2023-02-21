package router

import (
	"io"
	"net/http"
	proximityChannel "order-service/internal/api/controllers/proximity-channel"
	proximity_channel "order-service/internal/api/controllers/proximity-channel"
	"order-service/internal/api/controllers/users"
	"order-service/internal/api/middlewares"
	"order-service/internal/pkg/config"
	http_err "order-service/pkg/http-err"
	"os"

	"order-service/internal/api/problems"

	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Setup(config *config.Configuration) *gin.Engine {
	ginEngine := createGinEngine(config)
	//userController := buildUserController()

	// Routes
	addSwaggerDocs(ginEngine, config)
	addMiddlewares(ginEngine, config)
	//addUserController(ginEngine, config, userController)
	addProximityChannelController(ginEngine, config, proximityChannel.NewProximityChannelControllerDelegate())
	if config.Server.LogFileEnabled {
		addLogToFile(config)
	}
	return ginEngine
}

func addUserController(ginEngine *gin.Engine, config *config.Configuration, userController *users.UsersControllerDelegate) {
	users.RegisterHandlers(ginEngine, userController)

	swagger, err := users.GetSwagger()
	if err != nil {
		os.Exit(1)
	}

	ginEngine.Use(middleware.OapiRequestValidator(swagger))
}

func addProximityChannelController(ginEngine *gin.Engine, config *config.Configuration, proximityChannelDelegate *proximityChannel.ProximityChannelControllerDelegate) {
	//Default handler
	//proximityChannel.RegisterHandlers(ginEngine, proximityChannelDelegate)

	errorHandler := func(c *gin.Context, err error, statusCode int) {
		problem := problems.NewProblem(statusCode, c.Request.URL.Path, http.StatusText(statusCode), err.Error())
		c.JSONP(statusCode, problem)
	}
	// miuddl := []proximity_channel.MiddlewareFunc{proximity_channel.MiddlewareFunc(middlewares.NoMethodHandler())}

	options := proximity_channel.GinServerOptions{BaseURL: "/proximity-channel", Middlewares: nil, ErrorHandler: errorHandler}
	proximityChannel.RegisterHandlersWithOptions(ginEngine, proximityChannelDelegate, options)
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

// func buildUserController() *users.UsersControllerDelegate {
// 	userRepository := persistence.GetUserRepository()

// 	userService := service.NewUserService(*userRepository)
// 	userController := users.NewUsersControllerDelegate(*userService)

// 	return userController
// }

func addMiddlewares(ginEngine *gin.Engine, configs *config.Configuration) {
	// app.Use(gin.Recovery())
	ginEngine.Use(middlewares.CORS())
	ginEngine.NoMethod(middlewares.NoMethodHandler())
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
