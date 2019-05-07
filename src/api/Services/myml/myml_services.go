package myml

import (
	"github.com/mercadolibre/taller-go/src/api/Domain/myml"
	"github.com/mercadolibre/taller-go/src/api/Utils/apierrors"
	"net/http"
	"sync"
)

func GetRespuestaFromApiReceiver(userID int64) (*myml.JsonSuma, *apierrors.ApiError) {

	var respuesta myml.JsonSuma
	var wg sync.WaitGroup
	cE := make(chan *apierrors.ApiError)

	user := &myml.User{ID: int(userID)}
	err := user.Get()
	respuesta.User = *user
	if err != nil {
		return nil, &apierrors.ApiError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}

	wg.Add(1)

	go func() {
		err = <-cE
		wg.Done()
	}()

	go func() {
		err = respuesta.Site.Get(user.SiteID)
		cE <- err
	}()
	go func() {
		err = respuesta.Category.Get(user.SiteID)
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
