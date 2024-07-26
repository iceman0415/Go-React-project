package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bogdanguranda/go-react-example/db"
	db_mocks "github.com/bogdanguranda/go-react-example/db/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func setupMocks(t *testing.T) (*gomock.Controller, *db_mocks.MockDB) {
	mockCtrl := gomock.NewController(t)
	mockDB := db_mocks.NewMockDB(mockCtrl)
	return mockCtrl, mockDB
}

func mapResponse(resp *http.Response) (*Response, error) {
	respBodyBytes, _ := ioutil.ReadAll(resp.Body)
	response := Response{}
	return &response, json.Unmarshal(respBodyBytes, &response)
}

func mapPersonResp(resp *http.Response) (*db.Person, error) {
	respBodyBytes, _ := ioutil.ReadAll(resp.Body)
	person := db.Person{}
	return &person, json.Unmarshal(respBodyBytes, &person)
}

func mapPersonArrResp(resp *http.Response) ([]*db.Person, error) {
	respBodyBytes, _ := ioutil.ReadAll(resp.Body)
	var persons []*db.Person
	return persons, json.Unmarshal(respBodyBytes, &persons)
}

func TestCreatePerson_ValidPayload_Return200(t *testing.T) {
	mockCtrl, mockDB := setupMocks(t)
	defer mockCtrl.Finish()
	defAPI := NewDefaultAPI(mockDB)

	mockDB.EXPECT().CreatePerson(gomock.Any()).Return(nil)

	regReq := db.Person{
		Name:  "John Doe",
		Email: "something@gmail.com",
	}
	reqBody, _ := json.Marshal(regReq)
	req := httptest.NewRequest("POST", "localhost:8080/app/people", bytes.NewReader(reqBody))
	w := httptest.NewRecorder()

	defAPI.CreatePerson(w, req)
	_, err := mapResponse(w.Result())

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
}

func TestCreatePerson_MissingEmail_Return400(t *testing.T) {
	mockCtrl, mockDB := setupMocks(t)
	defer mockCtrl.Finish()
	defAPI := NewDefaultAPI(mockDB)

	regReq := db.Person{
		Name: "John Doe",
	}
	reqBody, _ := json.Marshal(regReq)
	req := httptest.NewRequest("POST", "localhost:8080/app/people", bytes.NewReader(reqBody))
	w := httptest.NewRecorder()

	defAPI.CreatePerson(w, req)
	_, err := mapResponse(w.Result())

	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
}

func TestDeletePerson_ValidRequest_Return200(t *testing.T) {
	mockCtrl, mockDB := setupMocks(t)
	defer mockCtrl.Finish()
	defAPI := NewDefaultAPI(mockDB)

	mockDB.EXPECT().DeletePerson("john@gmail.com").Return(nil)

	req := httptest.NewRequest("DELETE", "localhost:8080/app/people?email=john@gmail.com", nil)
	w := httptest.NewRecorder()

	defAPI.DeletePerson(w, req)
	_, err := mapResponse(w.Result())

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
}

func TestDeletePerson_MissingEmail_Return400(t *testing.T) {
	mockCtrl, mockDB := setupMocks(t)
	defer mockCtrl.Finish()
	defAPI := NewDefaultAPI(mockDB)

	req := httptest.NewRequest("DELETE", "localhost:8080/app/people", nil)
	w := httptest.NewRecorder()

	defAPI.DeletePerson(w, req)
	_, err := mapResponse(w.Result())

	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
}

func TestListPersons_ValidRequest_Return200(t *testing.T) {
	mockCtrl, mockDB := setupMocks(t)
	defer mockCtrl.Finish()
	defAPI := NewDefaultAPI(mockDB)

	persons := []*db.Person{
		&db.Person{
			Name:  "John Doe A",
			Email: "a.something@gmail.com",
		},
		&db.Person{
			Name:  "John Doe B",
			Email: "b.something@gmail.com",
		},
		&db.Person{
			Name:  "John Doe C",
			Email: "c.something@gmail.com",
		},
	}

	mockDB.EXPECT().ListPersons("name").Return(persons, nil)

	req := httptest.NewRequest("GET", "localhost:8080/app/people?orderBy=name", nil)
	w := httptest.NewRecorder()

	defAPI.ListPersons(w, req)
	personsResp, err := mapPersonArrResp(w.Result())
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, "John Doe A", personsResp[0].Name)
	assert.Equal(t, "John Doe B", personsResp[1].Name)
	assert.Equal(t, "John Doe C", personsResp[2].Name)
}

func TestListPersons_MissingOrderBy_DefaultToEmail_Return200(t *testing.T) {
	mockCtrl, mockDB := setupMocks(t)
	defer mockCtrl.Finish()
	defAPI := NewDefaultAPI(mockDB)

	persons := []*db.Person{
		&db.Person{
			Name:  "John Doe A",
			Email: "a.something@gmail.com",
		},
		&db.Person{
			Name:  "John Doe B",
			Email: "b.something@gmail.com",
		},
		&db.Person{
			Name:  "John Doe C",
			Email: "c.something@gmail.com",
		},
	}

	mockDB.EXPECT().ListPersons("email").Return(persons, nil)

	req := httptest.NewRequest("GET", "localhost:8080/app/people", nil)
	w := httptest.NewRecorder()

	defAPI.ListPersons(w, req)
	personsResp, err := mapPersonArrResp(w.Result())
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, "John Doe A", personsResp[0].Name)
	assert.Equal(t, "John Doe B", personsResp[1].Name)
	assert.Equal(t, "John Doe C", personsResp[2].Name)
}

func TestListPersons_OrderByWrongValue_Return400(t *testing.T) {
	mockCtrl, mockDB := setupMocks(t)
	defer mockCtrl.Finish()
	defAPI := NewDefaultAPI(mockDB)

	req := httptest.NewRequest("GET", "localhost:8080/app/people?orderBy=age", nil)
	w := httptest.NewRecorder()

	defAPI.ListPersons(w, req)
	_, err := mapResponse(w.Result())
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
}
