package mock

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MockCreateBooksPostJson(c *gin.Context, content []byte) {
	c.Request.Method = http.MethodPost
	c.Request.Header.Set("Content-Type", "application/json")

	c.Request.Body = io.NopCloser(bytes.NewBuffer(content))
}
