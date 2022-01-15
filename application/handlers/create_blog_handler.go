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
		post :=&entity.Blog{}
		err:= helpers.Decode(c,post) 
		if err != nil {
			response.JSON(c,"",http.StatusInternalServerError,nil,"Unable to decode"+err.Error())
		}
		
	}
}