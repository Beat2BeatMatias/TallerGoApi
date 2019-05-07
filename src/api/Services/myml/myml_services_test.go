package myml

import (
	"github.com/mercadolibre/taller-go/src/api/Domain/external_api"
	"testing"
)


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
