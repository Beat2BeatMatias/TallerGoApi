package myml

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/taller-go/src/api/Utils/apierrors"
	"io/ioutil"
	"net/http"
)

const (
	url = "https://api.mercadolibre.com/users/"
)

func (user *User) Get() *apierrors.ApiError {
	var data []byte

	if user.ID == 0 {
		return &apierrors.ApiError{
			Message: "UserID is empty",
			Status:  http.StatusInternalServerError,
		}
	}

	final := fmt.Sprintf("%s%d", url, user.ID)
	response, err := http.Get(final)
	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal(data, &user); err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return nil
}
