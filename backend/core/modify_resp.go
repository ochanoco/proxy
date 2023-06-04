package core

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainModifyResponse(proxy *OchanocoProxy, resp *http.Response) {
	fmt.Printf("=> %v\n", resp.Request.URL)
}

func LogModifyResponse(proxy *OchanocoProxy, res *http.Response, c *gin.Context) bool {
	_, err := logCommunication(res.Header, &res.Body, proxy)

	if err != nil {
		fmt.Printf("LogModifyResponse: %v\n", err)
		return false
	}

	return true
}
