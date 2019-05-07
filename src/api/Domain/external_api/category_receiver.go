package external_api

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/taller-go/src/api/Utils/apierrors"
	"io/ioutil"
	"net/http"
)

const urlSites = "https://api.mercadolibre.com/sites/"

func (category *Category) Get(siteID string, c chan apierrors.ApiError){
	var data []byte

	final := fmt.Sprintf("%s%s/categories", urlSites, siteID)
	response, err := http.Get(final)
	if err != nil {
		c<-apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	data, err = ioutil.ReadAll(response.Body)
	println(string(data))
	if err != nil {
		c<-apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal(data, &category); err != nil {
		c<-apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
}
