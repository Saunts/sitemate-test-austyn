package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Normally we put this in model folder, but since there's only 1 we just put it here instead
type issue struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

var tempDB = make(map[string]issue, 0)

const (
	urlWithID = "/issues/:id"

	issueNotExist = "Issue does not exist or already deleted"

	statusActive  = "Active"
	statusDeleted = "Deleted"
)

func main() {
	initTempDB()

	router := InitRouter()
	http.HandleFunc("/", homePageHandler)

	router.Run("localhost:8080")
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello world")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func initTempDB() {
	tempDB["1"] = issue{
		ID:          "1",
		Title:       "Button is misaligned",
		Description: "Button is misaligned",
		Type:        "UI",
		Status:      statusActive,
	}

	tempDB["2"] = issue{
		ID:          "2",
		Title:       "Font is wrong when on mobile",
		Description: "-",
		Type:        "UI",
		Status:      statusActive,
	}
}

func getIssues(c *gin.Context) {
	response := convertDataToResponse()

	c.IndentedJSON(http.StatusOK, response)
}

func getIssuesByID(c *gin.Context) {
	id := c.Param("id")

	issue, isExist := tempDB[id]
	if !isExist || issue.Status != statusActive {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": issueNotExist})
		return
	}

	c.IndentedJSON(http.StatusOK, issue)
}

func postIssues(c *gin.Context) {
	reqIssue, err := getAndValidateData(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	newIssue := issue{
		ID:          strconv.Itoa(len(tempDB) + 1),
		Title:       reqIssue.Title,
		Description: reqIssue.Description,
		Type:        reqIssue.Type,
		Status:      statusActive,
	}

	tempDB[newIssue.ID] = newIssue

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Issue successfully created with ID " + newIssue.ID})
}

func deleteIssues(c *gin.Context) {
	id := c.Param("id")

	issue, isExist := tempDB[id]
	if !isExist || issue.Status != statusActive {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": issueNotExist})
		return
	}

	issue.Status = statusDeleted
	tempDB[id] = issue

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Issue " + issue.Title + " with ID " + id + " is successfully deleted"})
}

func updateIssues(c *gin.Context) {
	id := c.Param("id")

	reqIssue, err := getAndValidateData(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	issue, isExist := tempDB[id]
	if !isExist || issue.Status != statusActive {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": issueNotExist})
		return
	}

	issue.Title = reqIssue.Title
	issue.Description = reqIssue.Description
	issue.Type = reqIssue.Type

	tempDB[id] = issue

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Issue " + issue.Title + " with ID " + id + " is successfully updated"})
}

func getAndValidateData(c *gin.Context) (issue, error) {
	reqIssue := issue{}
	c.BindJSON(&reqIssue)

	if reqIssue.Title == "" {
		return reqIssue, errors.New("title can't be empty")
	}

	if reqIssue.Description == "" {
		return reqIssue, errors.New("description can't be empty")
	}

	if reqIssue.Type == "" {
		return reqIssue, errors.New("type can't be empty")
	}

	return reqIssue, nil
}

func convertDataToResponse() []issue {
	response := []issue{}

	for _, data := range tempDB {
		if data.Status != statusActive {
			continue
		}

		response = append(response, data)
	}

	return response
}
