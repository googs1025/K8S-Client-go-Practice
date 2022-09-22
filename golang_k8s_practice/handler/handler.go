package handler

import (
	"github.com/gin-gonic/gin"
	"k8s_practice_client_go/golang_k8s_practice/common/api"
	"net/http"
)

func Success(c *gin.Context, data interface{}) {
	result := api.OK
	result.Data = data
	sendResponse(c, &result)
}

func Fail(c *gin.Context, err error) {

	result := api.InternalErr.FindMsg(err.Error())
	sendResponse(c, result)

}

func sendResponse(c *gin.Context, result *api.Result) {

	if result == nil || result.HTTPCode == 0 {
		c.JSON(http.StatusInternalServerError, api.InternalErr)
		return
	}
	c.JSON(result.HTTPCode, result)

}

func Response(c *gin.Context, data interface{}, err error) {
	if err != nil {
		Fail(c, err)
	} else {
		Success(c, data)
	}
}


