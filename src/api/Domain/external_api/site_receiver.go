package external_api

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/taller-go/src/api/Utils/apierrors"
	"io/ioutil"
	"net/http"
)

func (site *Site) Get(siteID string) *apierrors.ApiError {

	var data []byte

	final := fmt.Sprintf("%s%s", urlSites, siteID)
	response, err := http.Get(final)
	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	data, err = ioutil.ReadAll(response.Body)
	println(string(data))
	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal(data, &site); err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return nil
}