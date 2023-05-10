package router

import (
	"io"
	"net/http"
	"order-service/internal/api/controllers/order"
	"order-service/internal/api/middlewares"
	"order-service/internal/api/problems"
	"order-service/internal/pkg/config"
	"order-service/internal/pkg/persistence"
	"order-service/internal/pkg/service"
	http_err "order-service/pkg/http-err"
	"os"

	"github.com/gin-gonic/gin"
)

func Setup(config *config.Configuration) *gin.Engine {
	ginEngine := createGinEngine(config)
	// Routes
	// addSwaggerDocs(ginEngine, config)
	addMiddlewares(ginEngine, config)
	addOrderRestController(ginEngine, config, buildOrderDelegate())
	if config.Server.LogFileEnabled {
		addLogToFile(config)
	}
	return ginEngine
}

func addOrderRestController(ginEngine *gin.Engine, config *config.Configuration, proximityChannelDelegate *order.OrderControllerDelegate) {
	//Default handler
	//proximityChannel.RegisterHandlers(ginEngine, proximityChannelDelegate)
	errorHandler := func(c *gin.Context, err error, statusCode int) {
		problem := problems.NewProblem(statusCode, c.Request.URL.Path, http.StatusText(statusCode), err.Error())
		c.JSONP(statusCode, problem)
	}
	options := order.GinServerOptions{BaseURL: "/api", Middlewares: nil, ErrorHandler: errorHandler}
	order.RegisterHandlersWithOptions(ginEngine, proximityChannelDelegate, options)
}

func createGinEngine(config *config.Configuration) *gin.Engine {
	gin.SetMode(config.Server.Mode)
	if gin.Mode() == "release" {
		gin.DefaultWriter = io.Discard
	}
	ginEngine := gin.New()
	return ginEngine
}

// 	config := &ginSwagger.Config{
// 		URL: configs.Server.Api.DocsUrl,
// 	}
// 	r.GET(configs.Server.Api.SwaggerDocsUrl, ginSwagger.CustomWrapHandler(config, swaggerFiles.Handler))
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

func buildOrderDelegate() *order.OrderControllerDelegate {
	orderRepository := persistence.GetOrderRepository()

	orderService := service.NewOrderService(*orderRepository)
	orderControllerDelegate := order.NewOrderControllerDelegate(*orderService)

	return orderControllerDelegate
}
