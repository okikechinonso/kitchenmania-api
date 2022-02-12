package handlers

import (
	"kitchenmaniaapi/domain/entity"
	"kitchenmaniaapi/interfaces/helpers"
	"kitchenmaniaapi/interfaces/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *App) UpdatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		userI, exists := c.Get("user")
		if !exists {
			response.JSON(c, "can't get user from context", http.StatusInternalServerError, nil, "no user found")
			return
		}
		user := userI.(entity.User)
		_,err := app.DB.FindUserByEmail(user.Email)
		
		blog := entity.Blog{}
		err := helpers.Decode(c, blog)
		if err != nil{
			
		}
	}
}
