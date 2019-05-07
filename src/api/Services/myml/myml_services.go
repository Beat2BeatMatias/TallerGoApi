package myml

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/taller-go/src/api/Domain/myml"
	"github.com/mercadolibre/taller-go/src/api/Utils/apierrors"
	"io/ioutil"
	"net/http"
	"sync"
)

func GetUser(url string) myml.User {

	response, err := http.Get(url)

	var data []byte
	var user myml.User

	if err != nil {
		fmt.Println("Ocurrió un error %s\n", err)
	} else {
		data, _ = ioutil.ReadAll(response.Body)
		//fmt.Println(string(data))
	}

	err = json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("error:", err)
	}
	return user
}

const url = "https://api.mercadolibre.com/users/"

func GetUserFromApi(userID int64) (*myml.User, *apierrors.ApiError) {

	var data []byte
	var user myml.User

	if userID == 0 {
		return nil, &apierrors.ApiError{
			Message: "UserID is empty",
			Status:  http.StatusInternalServerError,
		}
	}

	final := fmt.Sprintf("%s%d", url, userID)
	response, err := http.Get(final)

	if err != nil {
		return nil, &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal(data, &user); err != nil {
		return nil, &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &user, nil
}

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
