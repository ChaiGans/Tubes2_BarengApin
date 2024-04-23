package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Input struct {
	StartTitleLink string `json:"StartTitleLink"`
	GoalTitleLink  string `json:"GoalTitleLink"`
	AlgoChoice int    `json:"AlgoChoice"`
}

type Result struct {
	Status      string   `json:"status"`
	Message     string   `json:"message"`
	ShortestPath []string `json:"shortestPath,omitempty"`
	ExecTime int64 `json:"exectime,omitempty"`
	NumChecked int `json:"numchecked,omitempty"`
}

type MultiResult struct {
	Status      string   `json:"status"`
	Message     string   `json:"message"`
	PossibleShortestPath [][]string `json:"possibleShortestPath,omitempty"`
	ExecTime int64 `json:"exectime,omitempty"`
	NumChecked int `json:"numchecked,omitempty"`
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/", postInformation)
	router.POST("/multi-api", postInformationMulti)
	router.Run(":8080")
}

func postInformation(c *gin.Context) {
	var userInput Input

	if err := c.BindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	result := processInputSingle(userInput)

	// Return the result as JSON.
	c.JSON(http.StatusOK, result)
}

func postInformationMulti(c *gin.Context) {
	var userInput Input

	if err := c.BindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	result := processInputMultiple(userInput)

	// Return the result as JSON.
	c.JSON(http.StatusOK, result)
}

func processInputSingle(input Input) Result {
	var err error

	switch input.AlgoChoice {
	case 1:
		start := time.Now()
		result, err, num_checked := IDS(input.StartTitleLink, input.GoalTitleLink)
		elapsed := time.Since(start).Milliseconds()
		if err == nil && result != nil {
			return Result{
				Status:      "Success",
				Message:     fmt.Sprintf("Processed using Algo IDS with StartTitle: %s and GoalTitle: %s", input.StartTitleLink, input.GoalTitleLink),
				ShortestPath: result,
				ExecTime: elapsed,
				NumChecked: num_checked,
			}
		}
	case 2:
		result, _, num_checked, _, duration, err := bfs(input.StartTitleLink, input.GoalTitleLink)
		if err == nil && result != nil {
			return Result{
				Status:      "Success",
				Message:     fmt.Sprintf("Processed using Algo 2 with StartTitle: %s and GoalTitle: %s", input.StartTitleLink, input.GoalTitleLink),
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

func processInputMultiple(input Input) MultiResult {
	var err error

	switch input.AlgoChoice {
	case 1:
		result, err, num_checked, exectime := MultiIDS(input.StartTitleLink, input.GoalTitleLink)
		if err == nil && result != nil {
			return MultiResult{
				Status:      "Success",
				Message:     fmt.Sprintf("Processed using Algo IDS with StartTitle: %s and GoalTitle: %s", input.StartTitleLink, input.GoalTitleLink),
				PossibleShortestPath: result,
				ExecTime: exectime,
				NumChecked: num_checked,
			}
		}
	case 2:
		result, err, num_checked, exectime := MultiIDS(input.StartTitleLink, input.GoalTitleLink)
		if err == nil && result != nil {
			return MultiResult{
				Status:      "Success",
				Message:     fmt.Sprintf("Processed using Algo 2 with StartTitle: %s and GoalTitle: %s", input.StartTitleLink, input.GoalTitleLink),
				PossibleShortestPath: result,
				ExecTime: exectime,
				NumChecked: num_checked,
			}
		}
	default:
		return MultiResult{Status: "Error", Message: "Invalid algorithm choice"}
	}

	if err != nil {
		return MultiResult{Status: "Error", Message: fmt.Sprintf("Error processing request: %v", err)}
	}
	return MultiResult{Status: "Error", Message: "Path not found"}
}