package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"fmt"
)

type Todo struct {
	title       string `json:"title"`
	description string `json:"description"`
	completed   bool   `json:"completed"`
	height      int    `json:"height"`
}

var db = make(map[string]Todo)

func main() {
	r := gin.Default()

	r.GET("/todos/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		todos, ok := db[user]
		if ok {
			c.JSON(200, gin.H{"user": user, "todos": todos})
		} else {
			c.JSON(200, gin.H{"user": user, "status": "no value"})
		}
	})

	r.POST("/add/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		var todo Todo
		if c.BindJSON(&todo) == nil {
			fmt.Println(todo)
			db[user] = todo
			c.JSON(200, gin.H{"status": "ok"})
		} else {
			c.JSON(200, gin.H{"status": "error"})
		}
	})

	//r.Use(cors.Default())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

	r.Run()
}
