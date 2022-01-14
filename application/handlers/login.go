package handlers

import (
	"kitchenmaniaapi/interfaces/helpers"
	"kitchenmaniaapi/interfaces/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *App) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginDetail := &struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}{}
		err := helpers.Decode(c, &loginDetail)
		if err != nil {
			response.JSON(c, "", http.StatusBadRequest, loginDetail, err.Error())
			return
		}

		user, err := a.DB.FindUserByEmail(loginDetail.Email)
		if err != nil {
			response.JSON(c, "", http.StatusBadRequest, nil, err.Error())
			return
		}
		log.Println(user)
		err = helpers.CompareHashAndPassword([]byte(loginDetail.Password), user.HashedPassword)
		if err != nil {
			response.JSON(c, "invalid password or email", http.StatusUnauthorized, nil, err.Error())
			return
		}
		user.HashedPassword = ""
		response.JSON(c, "login successful", http.StatusOK, gin.H{"user": user}, "")
	}
}
