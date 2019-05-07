package myml

import (
	"github.com/mercadolibre/taller-go/src/api/Domain/external_api"
	"github.com/mercadolibre/taller-go/src/api/Domain/myml"
	"github.com/mercadolibre/taller-go/src/api/Utils/apierrors"
	"sync"
	"testing"
)

func TestGetRespuestaFromApiReceiver(t *testing.T) {
	var respuesta myml.JsonSuma
	var wg sync.WaitGroup
	cE := make(chan *apierrors.ApiError)

	user := &external_api.User{ID: 12345678}
	err := user.Get()
	if err != nil{
		t.Error("No pudo conectarse. No se espera este error")
	}
	respuesta.User = *user

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

	if err != nil{
		t.Error("Error en la respuesta de la api consumida")
	}
}
func TestGetRespuestaFromApiReceiverNil(t *testing.T) {
	user := &external_api.User{SiteID: "dkajsdjk"}
	err := user.Get()
	if err == nil{
		t.Error("No pudo conectarse. Se espera este error")
	}
}
func TestGetRespuesta(t *testing.T) {
	respuesta, apiErr := GetRespuestaFromApiReceiver(12345678)
	if apiErr != nil{
		t.Error("No deberia tener este error")
	}
	if respuesta == nil{
		t.Error("No deberia tener este error")
	}
}
