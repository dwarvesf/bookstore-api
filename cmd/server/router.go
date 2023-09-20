package main

import (
	"github.com/dwarvesf/bookstore-api/docs"
	"github.com/dwarvesf/bookstore-api/pkg/handler"
	"github.com/dwarvesf/bookstore-api/pkg/handler/v1/portal"
	"github.com/dwarvesf/bookstore-api/pkg/logger/monitor"
	"github.com/dwarvesf/bookstore-api/pkg/middleware"
	"github.com/dwarvesf/bookstore-api/pkg/realtime"
	"github.com/dwarvesf/bookstore-api/pkg/service/jwthelper"
	"github.com/dwarvesf/bookstore-api/pkg/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func setupRouter(a App) *gin.Engine {
	docs.SwaggerInfo.Title = "Swagger API"
	docs.SwaggerInfo.Description = "This is a swagger for API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = a.cfg.BaseURL

	if a.cfg.Env == "local" {
		docs.SwaggerInfo.Schemes = []string{"http"}
	} else {
		docs.SwaggerInfo.Schemes = []string{"https"}
	}

	r := gin.New()

	r.Use(cors.New(
		cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
			AllowHeaders: []string{"Origin", "Host",
				"Content-Type", "Content-Length",
				"Accept-Encoding", "Accept-Language", "Accept",
				"X-CSRF-Token", "Authorization", "X-Requested-With", "X-Access-Token"},
			ExposeHeaders:    []string{"MeAllowMethodsntent-Length"},
			AllowCredentials: true,
		},
	))

	if a.cfg.SentryDSN != "" {
		// Once it's done, you can attach the handler as one of your middleware
		r.Use(otelgin.Middleware("api-service"))

		// recover when panic middleware
		r.Use(monitor.SentryPanicMiddleware(a.l))
	}

	// handlers
	publicHandler(r, a)
	authenticatedHandler(r, a)

	return r
}

func publicHandler(r *gin.Engine, a App) {
	h := handler.New(*a.cfg, a.monitor)
	portalHandler := portal.New(*a.cfg, a.l, a.repo, a.service, a.monitor)

	r.GET("/healthz", h.Healthz)

	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// api/v1
	apiV1 := r.Group("/api/v1")
	apiV1.POST("/auth/login", portalHandler.Login)
	apiV1.POST("/auth/signup", portalHandler.Signup)

	apiV1.GET("/sse", realtime.SSEHeadersMiddleware(), func(c *gin.Context) {
		u, err := a.realtimeServer.HandleConnection(c)
		if err != nil {
			a.l.Error(err, "failed to handle connection")
			util.HandleError(c, err)
			return
		}

		a.l.Infof("user %s connected", u.ID)
		a.realtimeServer.HandleEvent(c, *u, func(ginCtx *gin.Context, data any) error {
			a.l.Infof("data received: %v", data)
			return nil
		})
	})
}

func authenticatedHandler(r *gin.Engine, a App) {

	// api/v1
	authMw := middleware.NewAuthMiddleware(jwthelper.NewHelper(a.cfg.SecretKey))
	apiV1 := r.Group("/api/v1")
	apiV1.Use(authMw.WithAuth)
	portalHandler := portal.New(*a.cfg, a.l, a.repo, a.service, a.monitor)
	apiV1.GET("/me", portalHandler.Me)
	apiV1.PUT("/users", portalHandler.UpdateUser)
	apiV1.PUT("/users/password", portalHandler.UpdatePassword)
	apiV1.GET("/topics", portalHandler.GetTopics)

	bookGroup := apiV1.Group("/books")
	{
		bookGroup.GET("/", portalHandler.GetBooks)
		bookGroup.POST("/", portalHandler.CreateBook)
		bookGroup.PUT("/:id", portalHandler.UpdateBook)
		bookGroup.DELETE("/:id", portalHandler.DeleteBook)
	}
}
