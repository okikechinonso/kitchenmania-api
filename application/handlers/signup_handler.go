package server

import (
	"fmt"
	"kitchenmaniaapi/domain/entity"
	"kitchenmaniaapi/interfaces/helpers"
	"kitchenmaniaapi/interfaces/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func(s *Server) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *entity.User
		err := helpers.Decode(c, user)
		if err != nil {
			fmt.Errorf("enter all required field")
			return
		}

		hashedPassword,err := helpers.GenerateHashPassword(user.Password)
		if err != nil {
			fmt.Errorf("internal serve error",err)
			return
		}
		user.HashedPassword = string(hashedPassword)
		 _,err = s.DB.FindUserByEmail(user.Email)
		if err != nil {
			response.JSON(c,"", http.StatusNotFound,nil,"user email already exist")
			return
		}
		err = s.DB.NewUser(user)
		if err != nil {
			response.JSON(c, "", http.StatusInternalServerError, nil, "internal server error")
			return
		}
	}
}
