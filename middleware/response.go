package middleware

import (
	"bytes"
	"log"

	"github.com/gin-gonic/gin"
)

type responseLogger struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func newResponseLogger(w gin.ResponseWriter) *responseLogger {
	return &responseLogger{ResponseWriter: w, body: new(bytes.Buffer)}
}

func (r *responseLogger) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func (r *responseLogger) WriteString(s string) (int, error) {
	r.body.WriteString(s)
	return r.ResponseWriter.WriteString(s)
}

func ResponseLoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Wrap the original response writer
		wrappedWriter := newResponseLogger(c.Writer)
		c.Writer = wrappedWriter

		// Process the request
		c.Next()

		// After all subsequent handlers have executed, log the response
		statusCode := wrappedWriter.Status()
		responseBody := wrappedWriter.body.String() // Response body as a string
		log.Printf("Status Code: %d, Response Body: %s\n", statusCode, responseBody)
	}
}
