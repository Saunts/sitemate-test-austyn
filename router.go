package main

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/issues", getIssues)
	router.POST("/issues", postIssues)
	router.GET(urlWithID, getIssuesByID)
	router.DELETE(urlWithID, deleteIssues)
	router.PUT(urlWithID, updateIssues)

	return router
}
