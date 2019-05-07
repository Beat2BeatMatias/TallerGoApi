package myml

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/taller-go/src/api/Services/myml"
	"github.com/mercadolibre/taller-go/src/api/Utils/apierrors"
	"net/http"
	"strconv"
)

const (
	paramUserID="id"
)

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