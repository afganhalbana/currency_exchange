package responseutil

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// WriterOption for modifying the response
type WriterOption func(http.ResponseWriter)

// WithHeader add header to the response
func WithHeader(header string, value string) WriterOption {
	return func(w http.ResponseWriter) {
		w.Header().Set(header, value)
	}
}

// WithStatus define the response status code
// NOTE: need to be passed as the last argument because this function calls WriteHeader
func WithStatus(code int) WriterOption {
	return func(w http.ResponseWriter) {
		w.WriteHeader(code)
	}
}

// BadRequest response
func BadRequest(c *gin.Context, err string) {
	c.JSON(http.StatusBadRequest, gin.H{"message": err})
}

// NotFound response
func NotFound(w http.ResponseWriter) {
	http.Error(w, "", http.StatusNotFound)
}

// InternalServerError response
func InternalServerError(c *gin.Context, err string) {
	c.JSON(http.StatusInternalServerError, gin.H{"message": err})
}

// OK response
func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"data": data})
}
