package main

import (
	"bytes"
	"io/ioutil"
	"ka-auth-service/internal/application/service"
	"ka-auth-service/internal/infrastructure/db"
	"ka-auth-service/internal/infrastructure/env"
	"ka-auth-service/internal/infrastructure/repository"
	"ka-auth-service/internal/interfaces/controller"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)
}

func VerboseLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// generate a unique request ID
		requestID := uuid.New().String()

		// log request details
		log.WithFields(logrus.Fields{
			"request_id": requestID,
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
		}).Info("Incoming request")

		// log query parameters
		for key, values := range c.Request.URL.Query() {
			for _, value := range values {
				log.WithFields(logrus.Fields{
					"request_id":  requestID,
					"query_param": key,
					"value":       value,
				}).Info("Query parameter")
			}
		}

		// log path parameters
		for _, param := range c.Params {
			log.WithFields(logrus.Fields{
				"request_id": requestID,
				"path_param": param.Key,
				"value":      param.Value,
			}).Info("Path parameter")
		}

		// log request body (if present)
		if c.Request.Body != nil {
			bodyBytes, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				log.WithFields(logrus.Fields{
					"request_id": requestID,
					"error":      err,
				}).Error("Error reading request body")
			} else {
				log.WithFields(logrus.Fields{
					"request_id":   requestID,
					"request_body": string(bodyBytes),
				}).Info("Request body")
				// restore the request body for further handling by gin
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}

		// capture start time
		startTime := time.Now()

		// process request and handle panics
		defer func() {
			if r := recover(); r != nil {
				log.WithFields(logrus.Fields{
					"request_id":  requestID,
					"panic":       r,
					"stack_trace": string(debug.Stack()),
				}).Error("Panic occured")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			}
		}()

		// process request
		c.Next()

		// calculate the duration
		duration := time.Since(startTime)

		// log response status and time taken
		statusCode := c.Writer.Status()
		log.WithFields(logrus.Fields{
			"request_id":  requestID,
			"status_code": statusCode,
			"duration":    duration,
		}).Info("Response sent")

		// log detailed error information and stack trace for client or server errors
		if statusCode >= 400 && statusCode < 600 {
			log.WithFields(logrus.Fields{
				"request_id":    requestID,
				"status_code":   statusCode,
				"error_message": c.Errors.ByType(gin.ErrorTypePrivate).String(),
				"stack_trace":   string(debug.Stack()),
			}).Error("Error response")
		}
	}
}

func main() {
	config := env.LoadConfig()
	dbConn := db.NewDBConnection(config)

	userRepo := repository.NewUserRepoImpl(dbConn)
	userSessionRepo := repository.NewUserSessionRepoImpl(dbConn)
	employeeRepo := repository.NewEmployeeRepoImpl(dbConn)
	authService := service.NewAuthService(userRepo, userSessionRepo, employeeRepo)
	authController := controller.NewAuthController(authService)

	r := gin.New()
	r.Use(VerboseLogger())

	r.POST("/api/v1/auth", authController.Authentication)

	r.Run(":3000")
}
