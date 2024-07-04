package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var Posts = []Post{
	{ID: 1, Title: "Judul Postingan Pertama", Content: "Ini adalah postingan pertama di blog ini.", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 2, Title: "Judul Postingan Kedua", Content: "Ini adalah postingan kedua di blog ini.", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []User{
	{Username: "user1", Password: "pass1"},
	{Username: "user2", Password: "pass2"},
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()

		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		for _, user := range users {
			if user.Username == username && user.Password == password {
				return
			}
		}

		c.AbortWithStatus(http.StatusUnauthorized)
	} // TODO: replace this
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//Set up authentication middleware here // TODO: replace this
	r.Use(authMiddleware())

	r.GET("/posts", func(c *gin.Context) {

		if c.Query("id") == "" {
			c.JSON(http.StatusOK, gin.H{
				"posts": Posts,
			})
			return
		}

		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": "ID harus berupa angka",
			})
			return
		}

		for _, value := range Posts {
			if value.ID == id {
				c.JSON(http.StatusOK, gin.H{
					"post": value,
				})
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Postingan tidak ditemukan",
		})
		// TODO: answer here
	})

	r.POST("/posts", func(c *gin.Context) {
		var post Post
		err := c.ShouldBindJSON(&post)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		id := len(Posts) + 1
		timeNow := time.Now()
		post.ID = id
		post.CreatedAt = timeNow
		post.UpdatedAt = timeNow
		Posts = append(Posts, post)

		c.JSON(http.StatusCreated, gin.H{
			"message": "Postingan Berhasil ditambahkan",
			"post":    post,
		})
	})

	return r
}

func main() {
	r := SetupRouter()

	r.Run(":8080")
}
