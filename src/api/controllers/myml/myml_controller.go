package myml

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	myml2 "github.com/mercadolibre/taller-go/src/api/Domain/myml"
	"github.com/mercadolibre/taller-go/src/api/Services/myml"
	"github.com/mercadolibre/taller-go/src/api/Utils/apierrors"
	"net/http"
	"strconv"
	"time"
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

	cSite := make(chan myml2.Site)
	cCategory := make(chan myml2.Category)
	cError := make(chan *apierrors.ApiError)
	cOk := make(chan bool)

	user, apiErr := myml.GetUserFromApiReceiver(id)
	if apiErr != nil{
		context.JSON(apiErr.Status,apiErr)
		return
	}
	//152581223
	var site myml2.Site
	var category myml2.Category
	var errorR *apierrors.ApiError

	go func() {
		site=<-cSite
		category=<-cCategory
		errorR=<-cError
		cOk<-true
	}()

	go func() {myml.GetSiteApi(user.SiteID,cSite,cError)}()
	go func() {myml.GetCategoryApi(user.SiteID,cCategory,cError)}()


	select {

		case <- time.After(500 * time.Millisecond):
			errorR=&apierrors.ApiError{
				Message:"Superó el tiempo de respuesta",
				Status:http.StatusInternalServerError,
			}
		case <-cOk:

		}


	if errorR != nil {
		context.JSON(errorR.Status,errorR)
		return
	}
	jsonTotal:=myml2.JsonSuma{
		User:*user,
		Site:site,
		Category:category,
	}
	context.JSON(http.StatusOK,jsonTotal)
}