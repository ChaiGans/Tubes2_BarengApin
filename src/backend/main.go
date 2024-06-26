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
	ShortestPath [][]string `json:"shortestPath,omitempty"`
	ExecTime int64 `json:"exectime,omitempty"`
	NumChecked int `json:"numchecked"`
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
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
	var array_result [][]string

	switch input.AlgoChoice {
	case 1:
		start := time.Now()
		result, err, num_checked := IDS(input.StartTitleLink, input.GoalTitleLink)
		array_result = append([][]string{}, result)
		elapsed := time.Since(start).Milliseconds()
		if err == nil && result != nil {
			return Result{
				Status:      "Success",
				Message:     fmt.Sprintf("Processed using Algo IDS with StartTitle: %s and GoalTitle: %s", input.StartTitleLink, input.GoalTitleLink),
				ShortestPath: array_result,
				ExecTime: elapsed,
				NumChecked: num_checked,
			}
		}
	case 2:
		result, _, num_checked,  duration, err := bfs(input.StartTitleLink, input.GoalTitleLink)
		array_result = append([][]string{}, result)
		if err == nil && result != nil {
			return Result{
				Status:      "Success",
				Message:     fmt.Sprintf("Processed using Algo 2 with StartTitle: %s and GoalTitle: %s", input.StartTitleLink, input.GoalTitleLink),
				ShortestPath: array_result,
				ExecTime: duration,
				NumChecked: num_checked,
			}
		}
	case 3:
		result, err, num_checked, exectime := MultiIDS(input.StartTitleLink, input.GoalTitleLink)
		if err == nil && result != nil {
			return Result{
				Status:      "Success",
				Message:     fmt.Sprintf("Processed using Algo Multi-IDS with StartTitle: %s and GoalTitle: %s", input.StartTitleLink, input.GoalTitleLink),
				ShortestPath: result,
				ExecTime: exectime,
				NumChecked: num_checked,
			}
		}
	case 4:
		result, _, num_checked, exectime, err := bfsMultiPath(input.StartTitleLink, input.GoalTitleLink)
		if err == nil && result != nil {
			return Result{
				Status:      "Success",
				Message:     fmt.Sprintf("Processed using Algo Multi-BFS with StartTitle: %s and GoalTitle: %s", input.StartTitleLink, input.GoalTitleLink),
				ShortestPath: result,
				ExecTime: exectime,
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