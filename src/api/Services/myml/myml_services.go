package myml

import (
	"github.com/mercadolibre/taller-go/src/api/Domain/external_api"
	"github.com/mercadolibre/taller-go/src/api/Domain/myml"
	"github.com/mercadolibre/taller-go/src/api/Utils/apierrors"
	"net/http"
	"sync"
)

func GetRespuestaFromApiReceiver(userID int64) (*myml.JsonSuma, *apierrors.ApiError) {

	var respuestaTemp myml.JsonSuma
	var respuesta myml.JsonSuma
	var wg sync.WaitGroup

	site := &external_api.Site{}
	category := &external_api.Category{}

	cE := make(chan *apierrors.ApiError)
	c := make(chan myml.JsonSuma)

	user := external_api.User{ID: int(userID)}
	err := user.Get()
	if err != nil {
		return nil, &apierrors.ApiError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}

	wg.Add(3)
	go func() {
		respuestaTemp=<-c
		if respuestaTemp.Site != nil {
			respuesta.Site=respuestaTemp.Site
			wg.Done()
			respuestaTemp=<-c
			respuesta.Category=respuestaTemp.Category
			wg.Done()
		} else {
			respuesta.Category=respuestaTemp.Category
			wg.Done()
			respuestaTemp=<-c
			respuesta.Site=respuestaTemp.Site
			wg.Done()
		}
		respuesta.User=user
		err = <-cE
		wg.Done()
	}()
	go func() {
		err = site.Get(user.SiteID)
		c<-myml.JsonSuma{Site:site}
		cE <- err
	}()
	go func() {
		err = category.Get(user.SiteID)
		c<-myml.JsonSuma{Category:category}
		cE <- err
	}()
	wg.Wait()

	if err != nil {
		return nil, &apierrors.ApiError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}

	return &respuesta, nil
}
