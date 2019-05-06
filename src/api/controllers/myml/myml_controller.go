package myml

import (
	"encoding/json"
	myml2 "github.com/mercadolibre/taller-go/src/api/Domain/myml"
	"github.com/mercadolibre/taller-go/src/api/Utils/apierrors"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/taller-go/src/api/Services/myml"

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

	user, apiErr := myml.GetUserFromApiReceiver(id)
	if apiErr != nil{
		context.JSON(apiErr.Status,apiErr)
		return
	}

	var site myml2.Site
	var category myml2.Category
	var errorR *apierrors.ApiError

	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		site=<-cSite
		wg.Done()
		category=<-cCategory
		wg.Done()
		errorR=<-cError
		wg.Done()

	}()
	go func() {myml.GetSiteApi(user.SiteID,cSite,cError)}()
	go func() {myml.GetCategoryApi(user.SiteID,cCategory,cError)}()
	wg.Wait()

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