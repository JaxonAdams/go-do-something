package routes

import (
	"fmt"
	"go-do-something/database"
	"go-do-something/util"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

type NewTodoReq struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	DueDate     string `json:"due_date" binding: "required"`
}

type TodoListItem struct {
	UserID      string `json:"user_id"`
	TodoID      string `json:"todo_id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status"`
	DueDate     string `json:"due_date"`
	CreatedAt   string `json:"created_at"`
}

func RegisterTodoRoutes(r *gin.RouterGroup, svc *dynamodb.Client) {
	r.GET("/todo", func(c *gin.Context) {
		userID := "jaxontest@example.com" // TODO: add actual authentication
		getTodoItemsByUserID(userID, svc, c)
	})

	r.POST("/todo", func(c *gin.Context) {
		userID := "jaxontest@example.com" // TODO: add actual authentication
		createTodoItem(userID, svc, c)
	})
}

func getTodoItemsByUserID(userID string, svc *dynamodb.Client, c *gin.Context) {
	todoItems, err := database.GetUserTodoList(svc, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch to-do items"})
		return
	}

	todoList := make([]TodoListItem, 0, len(todoItems))

	for _, item := range todoItems {
		todoList = append(todoList, TodoListItem{
			UserID:      util.GetStringAttribute(item, "UserID"),
			TodoID:      util.GetStringAttribute(item, "TodoID"),
			Title:       util.GetStringAttribute(item, "Title"),
			Description: util.GetStringAttribute(item, "Description"),
			Status:      util.GetStringAttribute(item, "Status"),
			DueDate:     util.GetStringAttribute(item, "DueDate"),
			CreatedAt:   util.GetStringAttribute(item, "CreatedAt"),
		})
	}

	c.JSON(http.StatusOK, gin.H{"results": todoList})
}

func createTodoItem(userID string, svc *dynamodb.Client, c *gin.Context) {
	var json NewTodoReq
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.PutTodoItem(
		svc,
		userID,
		strconv.FormatInt(time.Now().UnixMilli(), 10),
		json.Title,
		json.Description,
		"pending",
		json.DueDate,
	)
	if err != nil {
		fmt.Printf("\nError creating new item:%v\n\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create new to-do item"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("Successfully created item: '%v'", json.Title)})
}
