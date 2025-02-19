package routes

import (
	"go-do-something/database"
	"go-do-something/util"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

type TodoListItem struct {
	UserID      string `json:"user_id"`
	TodoID      string `json:"todo_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	DueDate     string `json:"due_date"`
	CreatedAt   string `json:"created_at"`
}

func RegisterTodoRoutes(r *gin.RouterGroup, svc *dynamodb.Client) {
	r.GET("/todo", func(c *gin.Context) {
		userID := "jaxontest@example.com" // TODO: add actual authentication
		getTodoItemsByUserID(userID, svc, c)
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
