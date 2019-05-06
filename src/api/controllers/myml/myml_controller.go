package myml

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/taller-go/src/api/Services/myml"
	"github.com/mercadolibre/taller-go/src/api/Utils/apierrors"
	"net/http"
	"strconv"
)

const (
	paramUserID="id"
)

func UserData(context *gin.Context){

	userID := context.Param(paramUserID)
	id, err := strconv.ParseInt(userID,10,64)

	if err != nil {
		return
	}
	url := "https://api.mercadolibre.com/users/" + string(id)

	user := myml.GetUser(url)

	userString, _ :=json.Marshal(&user)

	context.String(200, string(userString))

}
func GetUserData(context *gin.Context){

	userID := context.Param(paramUserID)
	id, err := strconv.ParseInt(userID,10,64)

	if err != nil {
		apiErr := &apierrors.ApiError{
			Message:err.Error(),
			Status:http.StatusInternalServerError,
		}
		context.JSON(apiErr.Status,apiErr)
		return
	}

	user, apiErr := myml.GetUserFromApi(id)
	if apiErr != nil{
		context.JSON(apiErr.Status,apiErr)
		return
	}
	context.JSON(http.StatusOK,user)
}


func GetUserDataReceiver(context *gin.Context){
	userID := context.Param(paramUserID)
	id, err := strconv.ParseInt(userID,10,64)
	if err != nil {
		apiErr := &apierrors.ApiError{
			Message:err.Error(),
			Status:http.StatusInternalServerError,
		}
		context.JSON(apiErr.Status,apiErr)
		return
	}

	respuesta, apiErr := myml.GetRespuestaFromApiReceiver(id)
	if apiErr != nil{
		context.JSON(apiErr.Status,apiErr)
		return
	}

	context.JSON(http.StatusOK,respuesta)
}