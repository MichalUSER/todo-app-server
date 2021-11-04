package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	Height      int    `json:"height"`
}

var db = make(map[string][]Todo)

func main() {
	r := gin.Default()

	//r.Use(cors.Default())
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

	r.GET("/todos/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		todo, ok := db[user]
		if ok {
			c.JSON(200, gin.H{"user": user, "todos": todo})
		} else {
			c.JSON(200, gin.H{"user": user, "status": "no value"})
		}
	})

	r.POST("/add/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		var todos []Todo
		if err := c.BindJSON(&todos); err != nil {
			c.JSON(200, gin.H{"status": err.Error()})
		} else {
			db[user] = todos
			c.JSON(200, gin.H{"status": "ok"})
		}
	})

	r.Run()
}
