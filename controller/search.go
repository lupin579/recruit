package controller

import (
	"recruit/pkg/code"
	"recruit/service"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	key := c.Param("key")
	search := new(service.Search)
	search.Key = key
	searchResult, err := search.SearchService()
	if err != nil {
		ResponseError(c, code.GetError, err)
		return
	}
	ResponseSuccess(c, code.Success, searchResult)
}
