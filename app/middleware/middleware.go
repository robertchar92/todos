package middleware

import (
	"strconv"

	"todo/utils/errors"

	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stoewer/go-strcase"
	"gorm.io/gorm"
)

type Config struct {
	// put middleware config here
	Db *gorm.DB
}

type Middleware struct {
	config Config
}

func New(cfg Config) *Middleware {
	return &Middleware{
		config: cfg,
	}
}

func (m *Middleware) AuthHandle() gin.HandlerFunc {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Key:              []byte(os.Getenv("JWT_SECRET")),
		MaxRefresh:       time.Hour,
		TimeFunc:         time.Now,
		SigningAlgorithm: "HS512",
	})
	if err != nil {
		log.Fatal("jwt-error:" + err.Error())
	}
	return authMiddleware.MiddlewareFunc()
}

func (m *Middleware) VersionHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check version
		clientVersion, _ := strconv.ParseFloat(c.GetHeader("Version"), 64)
		requiredVersion, _ := strconv.ParseFloat(os.Getenv("APP_VERSION"), 64)

		if clientVersion < requiredVersion {
			err := errors.ErrBadRequest
			err.Message = "please update your apps to newer version"
			c.AbortWithStatusJSON(err.HTTPCode, gin.H{"errors": err.Message})
			return
		}

		c.Next()
	}
}

func (m *Middleware) ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// Only run if there are some errors to handle
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				// Find out what type of error it is

				switch e.Type {
				case gin.ErrorTypePublic:
					// Only output public errors if nothing has been written yet
					// if !c.Writer.Written() {
					// check if it is part of custom error
					// if err, ok := e.Err.(errors.CustomError); ok {
					// 	// report if stack traced
					// 	if stackTracer, ok := err.ErrWithStack.(errors.StackTracer); ok {
					// 		request_util.HandleExceptionLogging(c, err.HTTPCode, err.ErrWithStack, &stackTracer)
					// 	}

					// 	// print the underlying error and return the specified message to user
					// 	c.JSON(err.HTTPCode, gin.H{"errors": err.Message})
					// } else {
					// 	c.JSON(c.Writer.Status(), gin.H{"errors": e.Error()})
					// }
					// }
				case gin.ErrorTypeBind:
					errs, ok := e.Err.(validator.ValidationErrors)
					if ok {
						list := make(map[string]string)
						for _, err := range errs {
							list[strcase.SnakeCase(err.Field())] = validationErrorToText(err)
						}

						// Make sure we maintain the preset response status
						status := http.StatusUnprocessableEntity
						if c.Writer.Status() != http.StatusOK {
							status = c.Writer.Status()
						}
						c.JSON(status, gin.H{"errors": list})
					} else {
						c.JSON(422, gin.H{"errors": "please make sure to provide the correct data type or format"})
					}

				default:
					// Log all other errors
					//rollbar.RequestError(rollbar.ERR, c.Request, e.Err)
				}

			}
			// If there was no public or bind error, display default 500 message
			if !c.Writer.Written() {
				c.JSON(http.StatusInternalServerError, gin.H{"errors": "something went wrong"})
			}
		}
	}
}
