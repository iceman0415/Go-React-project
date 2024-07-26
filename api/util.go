package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/bogdanguranda/go-react-example/db"
)

// Response is the JSON response of the API when no information about persons is requested.
type Response struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

// mapPersonPayload maps a http request to create a new person.
func (dAPI *DefaultAPI) mapPersonPayload(r *http.Request) (*db.Person, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	person := db.Person{}
	if err := json.Unmarshal(body, &person); err != nil {
		return nil, err
	}

	return &person, dAPI.validatePerson(&person)
}

// validatePerson validates mandatory fields on a person data.
func (dAPI *DefaultAPI) validatePerson(person *db.Person) error {
	if person.Email == "" {
		return errors.New("Missing mandatory field 'email'")
	}

	return nil
}
