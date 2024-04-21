package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Input struct {
	StartTitle string `json:"StartTitle"`
	GoalTitle  string `json:"GoalTitle"`
	AlgoChoice int    `json:"AlgoChoice"`
}

type Result struct {
	Status      string   `json:"status"`
	Message     string   `json:"message"`
	ShortestPath []string `json:"shortestPath,omitempty"`
	ExecTime int64 `json:"exectime,omitempty"`
	NumChecked int `json:"numchecked,omitempty"`
}

func main() {
	router := gin.Default()
	router.POST("/", postInformation)
	router.Run(":8080")
}

func postInformation(c *gin.Context) {
	var userInput Input

	if err := c.BindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	result := processInput(userInput)

	// Return the result as JSON.
	c.JSON(http.StatusOK, result)
}

func processInput(input Input) Result {
	var err error

	var starting_title_link string = titleToUrl(input.StartTitle);
	var goal_title_link string = titleToUrl(input.GoalTitle);

	switch input.AlgoChoice {
	case 1:
		start := time.Now()
		result, err, num_checked := IDS(starting_title_link, goal_title_link)
		elapsed := time.Since(start).Milliseconds()
		if err == nil && result != nil {
			return Result{
				Status:      "Success",
				Message:     fmt.Sprintf("Processed using Algo IDS with StartTitle: %s and GoalTitle: %s", starting_title_link, goal_title_link),
				ShortestPath: result,
				ExecTime: elapsed,
				NumChecked: num_checked,
			}
		}
	case 2:
		result, _, num_checked, _, duration, err := bfs(starting_title_link, goal_title_link)
		if err == nil && result != nil {
			return Result{
				Status:      "Success",
				Message:     fmt.Sprintf("Processed using Algo 2 with StartTitle: %s and GoalTitle: %s", starting_title_link, goal_title_link),
				ShortestPath: result,
				ExecTime: duration,
				NumChecked: num_checked,
			}
		}
	default:
		return Result{Status: "Error", Message: "Invalid algorithm choice"}
	}

	if err != nil {
		return Result{Status: "Error", Message: fmt.Sprintf("Error processing request: %v", err)}
	}
	return Result{Status: "Error", Message: "Path not found"}
}

func titleToUrl(title string) string {
	return "https://en.wikipedia.org/wiki/" + strings.ReplaceAll(title, " ", "_")
}