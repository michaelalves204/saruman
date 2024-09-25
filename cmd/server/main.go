package main

import (
	"fmt"
	"net/http"
	"os"
	file "saruman/cmd/file"

	"github.com/gin-gonic/gin"
)

type Request struct {
	ProcessName string
	Message     string
}

func main() {
	router := gin.Default()
	router.GET("/health_check", health)
	router.POST("/processes", processMonitor)

	f, _ := os.ReadFile("saruman.txt")

	fmt.Printf(string(f))

	port := "8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	router.Run(fmt.Sprintf("localhost:%s", port))
}

func health(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}

func processMonitor(c *gin.Context) {
	var req Request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	processName := req.ProcessName
	message := req.Message

	go file.Create(processName, message)

	c.JSON(http.StatusOK, gin.H{
		"status": "enqueue",
	})
}
