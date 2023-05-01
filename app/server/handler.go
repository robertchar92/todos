package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"todo/app/middleware"
	"todo/utils/validators"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// All handler that need to be registered MUST implement this interface
type Handler interface {
	Register(r *gin.Engine, m *middleware.Middleware)
}

func BuildHandler(middleware *middleware.Middleware, handlers ...Handler) http.Handler {
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	// register all custom validator here
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("unique",
			validators.CustomValidator.Unique())
		if err != nil {
			log.Println("error when applying unique validator")
		}
		err = v.RegisterValidation("exist", validators.CustomValidator.Exist())
		if err != nil {
			log.Println("error when applying exist validator")
		}
		err = v.RegisterValidation("value", validators.CustomValidator.Value())
		if err != nil {
			log.Println("error when applying value validator")
		}
	}

	// apply global middleware here
	// config := cors.Config{
	// 	//AllowOrigins: strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ","),
	// 	//AllowWebSockets:  true,
	// 	//AllowCredentials: true,
	// 	AllowAllOrigins: true,
	// 	AllowWildcard:   true,
	// 	AllowHeaders: []string{
	// 		"Accept",
	// 		"Accept-Encoding",
	// 		"Bearer",
	// 		"Connection",
	// 		"Cookie",
	// 		"Content-Length",
	// 		"Content-Type",
	// 		"Origin",
	// 		"Authorization",
	// 		"X-Forwarded-For",
	// 		"X-Real-Ip",
	// 		"User-Agent",
	// 		"Lang",
	// 		"Version",
	// 		"Access-Control-Request-Method",
	// 		"Access-Control-Request-Headers",
	// 		"Access-Control-Allow-Origin",
	// 	},
	// 	AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
	// }

	// router.Use(cors.New(config))
	router.Use(middleware.ErrorHandle())

	// set max upload file size
	//router.MaxMultipartMemory = 8 << 20  // 8 MiB

	// router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	// router.GET("/healthz", healthz)
	// serverRoute := router.Group("/server", middleware.BasicHandle())
	// {
	// 	serverRoute.GET("/err_test", errTest)
	// }

	router.Use(static.Serve("/public", static.LocalFile("public", false)))
	// router.StaticFS("/assets", http.Dir("public"))
	//router.GET("/handling-midtrans", test)
	//router.GET("/test",test)

	// start registering routes from all handlers
	for _, reg := range handlers {
		reg.Register(router, middleware)
	}

	// 404 not found function
	router.NoRoute(notFound)

	return router
}

func healthz(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprint("ok:", runtime.NumGoroutine()))
}

// func errTest(c *gin.Context) {
// 	panic("Error Test")
// 	c.String(http.StatusOK, fmt.Sprint("ok:", runtime.NumGoroutine()))
// }

// func test(c *gin.Context) {
// 	x, _ := ioutil.ReadAll(c.Request.Body)
// 	log.Println("===========")
// 	log.Printf("%s", string(x))
// }

func notFound(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotFound)
}
