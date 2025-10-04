package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"myapp/types"
	"github.com/go-playground/validator/v10"
	"errors"
	"fmt"
)

func UserRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/user", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})

		api.POST("/user", func(c *gin.Context) {

			var validate *validator.Validate
			validate = validator.New()
			var user types.User
			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"errors": []string{"Dữ liệu không hợp lệ"}})
				return
			}

			if err := validate.Struct(user); err != nil {
				var ve validator.ValidationErrors
				if errors.As(err, &ve) {
					msgs := []string{}
					for _, fe := range ve {
						switch fe.Tag() {
						case "required":
							msgs = append(msgs, fmt.Sprintf("%s là bắt buộc", fe.Field()))
						case "email":
							msgs = append(msgs, "Email không hợp lệ")
						case "gte":
							msgs = append(msgs, "Tuổi phải lớn hơn hoặc bằng 18")
						case "lte":
							msgs = append(msgs, "Tuổi phải nhỏ hơn hoặc bằng 65")
						case "phonevn":
							msgs = append(msgs, "Số điện thoại không hợp lệ (phải 09 hoặc 08)")
						}
					}
					c.JSON(http.StatusBadRequest, gin.H{"errors": msgs})
					return
				}
			}

			c.JSON(http.StatusOK, gin.H{"message": "Đăng ký thành công"})
			})
	}
}
