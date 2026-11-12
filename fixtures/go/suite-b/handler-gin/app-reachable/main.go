package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	"example.com/cbench/handler-gin/liba"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/health", health)
	r.GET("/fingerprint", fingerprint)
	if len(os.Args) > 1 && os.Args[1] == "serve" {
		if err := r.Run("127.0.0.1:8080"); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}
	// Without "serve", answer one request in process and exit.
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/fingerprint?data=hello", nil))
	fmt.Println(rec.Code, rec.Body.String())
}

func health(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

func fingerprint(c *gin.Context) {
	c.String(http.StatusOK, "%x", liba.Fingerprint([]byte(c.Query("data"))))
}
