package main

import (
	"database/sql"
	"log"
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/adolfovega/gochallenge/dbutils"
)


// DB Driver used to query database
var DB *sql.DB

// TaskResource structure
type TaskResource struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Priority	int    `json:"priority"`
	Description	string `json:"description"`
	StartTime 	string `json:"start_time"`
	EndTime		string `json:"end_time"`
}

// GetTask returns the task detail
func GetTask(c *gin.Context) {
	var task TaskResource
	id := c.Param("task_id")
	err := DB.QueryRow("select id, name, priority, description, CAST(start_time as CHAR), CAST(end_time as CHAR) from tasks where id=$1;", id).Scan(&task.ID, &task.Name, &task.Priority, &task.Description, &task.StartTime, &task.EndTime)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error() + " id: "+ id,
		})
	} else {
		c.JSON(200, gin.H{
			"result": task,
		})
	}
}

// GetTasks returns all tasks
func GetTasks(c *gin.Context) {
	var tasks []TaskResource
	rows, err := DB.Query("select id, name, priority, description, start_time, end_time FROM tasks;")
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		for rows.Next() {
			var task TaskResource
			err = rows.Scan(&task.ID, &task.Name, &task.Priority, &task.Description, &task.StartTime, &task.EndTime)
			if err != nil {
				log.Println(err)
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
			}
			tasks = append(tasks, task)
		}

		c.JSON(200, gin.H{
			"result": tasks,
		})
	}
}

// CreateTask handles the POST
func CreateTask(c *gin.Context) {
	var task TaskResource
	// Parse the body into our resrource
	if err := c.BindJSON(&task); err == nil {
		// Format Time to Go time format
		statement, _ := DB.Prepare("insert into tasks (name, priority, description, start_time, end_time) values ($1, $2, $3, $4, $5)")
		result, _ := statement.Exec(task.Name, task.Priority, task.Description, task.StartTime, task.EndTime)
		if err == nil { 
			newID, _ := result.LastInsertId()
			task.ID = int(newID)
			c.JSON(http.StatusOK, gin.H{
				"result": task,
			})
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

// RemoveTask handles the DELETE
func RemoveTask(c *gin.Context) {
	id := c.Param("task_id")
	statement, _ := DB.Prepare("delete from tasks where id=$1")
	_, err := statement.Exec(id)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.String(http.StatusOK, "")
	}
}

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "P@ssw0rd"
  dbname   = "tododb"
)

func main() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)
	DB, err = sql.Open("postgres",psqlInfo)
	if err != nil {
		log.Println("Driver creation failed!")
	}
	dbutils.Initialize(DB)
	router := gin.Default()
	// Add routes to REST verbs
	router.GET("/v1/tasks/:task_id", GetTask)
	router.GET("/v1/tasks", GetTasks)
	router.POST("/v1/tasks", CreateTask)
	router.DELETE("/v1/tasks/:task_id", RemoveTask)

	router.Run(":8000") // Default listen and serve on 0.0.0.0:8000
}

