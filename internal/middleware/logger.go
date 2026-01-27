package middleware

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//going to log method, path, status, duration of each HTTP request

// Logger is a Gin middleware that logs each incoming HTTP request.
//
// What this middleware should do:
//
// 1) Before the request is processed:
//    - record the current time (start time)
//
// 2) Let Gin process the request normally:
//    - call c.Next()
//
// 3) After the request is processed:
//    - calculate how long the request took
//    - get information about the request:
//        * HTTP method
//        * request path
//        * response status code
//
// 4) Log these details in a readable format
//
// Example log output we want:
//
// [REQUEST] GET /finance -> 200 (12ms)
//
// This helps us understand:
// - which endpoints are being called
// - how fast they respond
// - if any of them return errors

func Logger() gin.HandlerFunc {

	// return a gin middleware function
	return func(c *gin.Context) {

		// STEP 1:
		// Save the start time of the request
		requestTimeStart := time.Now()

		// STEP 2:
		// Continue processing the request
		// (this calls the next middleware / handler)
		c.Next()

		// STEP 3:
		// After the request is done, calculate duration
		duration := time.Since(requestTimeStart)

		// STEP 4:
		// Read useful info from Gin context:
		// - HTTP method
		// - request path
		// - response status code
		method := c.Request.Method
		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}
		statusCode := c.Writer.Status()
		requestip := c.ClientIP()
		// STEP 5:
		// Log everything using log.Println or slog.Info
		logRequest(method, path, statusCode, duration, requestip)

	}
}

func logRequest(method, path string, statusCode int, duration time.Duration, requestip string) {
	// Format the log message
	logMessage := "[REQUEST] " + method + " " + path +
		" -> " + strconv.Itoa(statusCode) + " (" + duration.String() + ") from IP: " + requestip
	// Print the log message
	log.Println(logMessage)
}
