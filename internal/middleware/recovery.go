package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Recovery is a Gin middleware that catches panics and prevents the app from crashing.
//
// What this middleware do:
//
// 1) Wrap the request lifecycle in a defer function
// 2) If a panic happens anywhere in the request chain:
//    - catch it using recover()
//    - log the error
//    - return a JSON response with HTTP 500
//
// Example response:
//
// {
//   "error": "internal server error"
// }
//
// Why this is important?
// - Without recovery, a single panic can crash the whole server
// - With recovery, we isolate errors per request
func Recovery() gin.HandlerFunc {

	return func(c *gin.Context) {

		// STEP 1:
		// Use defer to catch panics
		defer func() {
			if err := recover(); err != nil {
				// STEP 2:
				// A panic happened!
				// Log the error
				log.Printf("Panic recovered: %v", err)
				// Return JSON error response
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				// Abort the request to prevent further processing
				c.Abort()
			}
		}()

		// STEP 3:
		// Continue processing the request
		c.Next()
		}
	}
