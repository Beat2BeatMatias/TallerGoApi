package myml

import (
	"github.com/mercadolibre/taller-go/src/api/Domain/external_api"
	"github.com/mercadolibre/taller-go/src/api/Domain/myml"
	"github.com/mercadolibre/taller-go/src/api/Utils/apierrors"
	"net/http"
)

func GetRespuestaFromApiReceiver(userID int64) (*myml.JsonSuma, *apierrors.ApiError) {

	var respuesta myml.JsonSuma

	site := external_api.Site{}
	category := external_api.Category{}

	cE := make(chan apierrors.ApiError)
	c := make(chan myml.JsonSuma)
	cOk := make(chan bool)

	user := external_api.User{ID: int(userID)}
	err := user.Get()
	if err != nil {
		return nil, &apierrors.ApiError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}

	go func() {
		respuesta=<-c
		respuesta=<-c
		respuesta.User=user
		cOk<-true
	}()

	go func() {
		site.Get(user.SiteID,cE)
		c<-myml.JsonSuma{Site:site}

	}()
	go func() {
		category.Get(user.SiteID,cE)
		c<-myml.JsonSuma{Category:category}
	}()

	select {

		case err := <- cE:
			return nil, &err

		case <- cOk:

	}

	return &respuesta, nil
}
