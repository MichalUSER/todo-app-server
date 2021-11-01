package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	Height      int    `json:"height"`
}
//type Todo struct {
//	Name string `json:"name"`
//}

var db = make(map[string]Todo)

func main() {
	r := gin.Default()

	//r.Use(cors.Default())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

	r.GET("/todos/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		todo, ok := db[user]
		if ok {
			c.JSON(200, gin.H{"user": user, "todo": todo})
		} else {
			c.JSON(200, gin.H{"user": user, "status": "no value"})
		}
	})

	r.POST("/add/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		var todo Todo
		if err := c.BindJSON(&todo); err != nil {
			c.JSON(200, gin.H{"status": err.Error()})
		} else {
			db[user] = todo
			c.JSON(200, gin.H{"status": "ok"})
		}
	})

	r.Run()
}
