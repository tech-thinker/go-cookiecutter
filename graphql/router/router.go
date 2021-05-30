package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tech-thinker/go-cookiecutter/api/controllers"
	"github.com/tech-thinker/go-cookiecutter/service/initializer"
)

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Init sets router
func Init(dependencies initializer.Services) *gin.Engine {
	router := gin.Default()

	// Add middlewares
	router.Use(ginlogrus.Logger(logger.Log))
	router.Use(middlewares.GinContextToContext())

	// setup cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = []string{"Accept", "Accept-CH", "Accept-Charset", "Accept-Datetime", "Accept-Encoding", "Accept-Ext", "Accept-Features", "Accept-Language", "Accept-Params", "Accept-Ranges", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Origin", "Access-Control-Expose-Headers", "Access-Control-Max-Age", "Access-Control-Request-Headers", "Access-Control-Request-Method", "Authorization", "Content-Type"}
	corsConfig.AllowAllOrigins = true

	// setup cors middleware
	router.Use(cors.New(corsConfig))

	// panic recovery
	router.Use(nice.Recovery(func(c *gin.Context, err interface{}) {
		var e error
		if err == nil {
			e = nil
		} else {
			e = err.(error)
		}
		logger.Log.WithError(e).Error(`GraphQL panic`)
	}))

	introspectionEnabled := true
	if config.BuildEnv() != "dev" {
		router.Use(bugsnaggin.AutoNotify(instance.BugSnagConfig()))
		router.Use(middlewares.Bugsnag())
		if config.BuildEnv() == "production" {
			introspectionEnabled = false
		}
	}

	router.NoRoute(func(c *gin.Context) {
		utils.HandleError(c, constants.NotFound, errors.New("not found"))
	})

	pingHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	}

	router.GET("/ping", pingHandler)

	router.POST("/query", middlewares.Auth(dependencies.AuthRepo()), graphqlHandler(dependencies, introspectionEnabled))

	if config.BuildEnv() != "production" {
		router.GET("/", playgroundHandler())
	}

	return router
}
