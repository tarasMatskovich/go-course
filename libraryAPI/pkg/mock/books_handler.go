package mock

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MockCreateBooksPostJson(c *gin.Context, content interface{}) {
	c.Request.Method = http.MethodPost
	c.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		log.Fatal(err)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}
