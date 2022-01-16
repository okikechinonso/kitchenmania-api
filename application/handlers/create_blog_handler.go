package handlers

import (
	"kitchenmaniaapi/domain/entity"
	"kitchenmaniaapi/interfaces/helpers"
	"kitchenmaniaapi/interfaces/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *App) CreateBlog() gin.HandlerFunc{
	return func(c *gin.Context){
		userI,exist := c.Get("user");
		if  !exist {
			response.JSON(c,"can't get user from context",http.StatusInternalServerError,nil,"no user found")
			return
		}
		user,ok := userI.(entity.User)
		if !ok {
			response.JSON(c,"",http.StatusInternalServerError,nil,"Unable to decode")
			return
		}
		post :=&entity.Blog{}
		err:= helpers.Decode(c,post) 
		if err != nil {
			response.JSON(c,"",http.StatusInternalServerError,nil,"Unable to decode"+err.Error())
			return
		}
		
		post.Title = c.PostForm("title")
		post.Body  = c.PostForm("body")
		post.Author = user.FirstName + user.LastName

		form, err := c.MultipartForm()

		
	}
}