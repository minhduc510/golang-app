package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"myapp/types"
	"myapp/helpers"
)

func UserRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/user", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})

		api.POST("/user", func(c *gin.Context) {

			user := types.RegisterUser{
				Phone:    c.PostForm("phone"),
				Email:    c.PostForm("email"),
				Username: c.PostForm("username"),
			}

			if user.Phone == "" || user.Email == "" || user.Username == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
				return
			}

			if !helpers.ValidatePhone(user.Phone) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone number format"})
				return
			}

			if !helpers.ValidateEmail(user.Email) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
				return
			}

			if !helpers.ValidateLength(user.Username, 3, 20) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Username must be between 3 and 20 characters"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"data": user})
		})
	}
}
