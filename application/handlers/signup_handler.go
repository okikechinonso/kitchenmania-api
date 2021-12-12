package handlers

import (
	"fmt"
	"kitchenmaniaapi/domain/entity"
	"kitchenmaniaapi/interfaces/helpers"
	"kitchenmaniaapi/interfaces/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *App) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user entity.User
		log.Println(c.Request.Body);
		err := helpers.Decode(c, &user)
		log.Println(user)
		if err != nil {
			// fmt.Errorf("enter all required field")
			response.JSON(c, "", http.StatusBadRequest, nil, "enter all field")
			log.Printf("Error: %v", err.Error())
			return
		}
		fmt.Println(user)
		hashedPassword, err := helpers.GenerateHashPassword(user.Password)
		if err != nil {
			fmt.Errorf("internal serve error", err)
			return
		}
		user.HashedPassword = string(hashedPassword)
		_, err = s.DB.FindUserByEmail(user.Email)
		if err != nil {
			response.JSON(c, "", http.StatusNotFound, nil, "user email already exist")
			return
		}
		err = s.DB.NewUser(&user)
		if err != nil {
			response.JSON(c, "", http.StatusInternalServerError, nil, "internal server error")
			return
		}
	}
}
