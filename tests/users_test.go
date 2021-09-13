package tests

import (
	"../controllers"
	"../endpoints"
	"../models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func assertStatusCode(t *testing.T, response *httptest.ResponseRecorder, wantStatusCode int) {
	if response.Result().StatusCode != wantStatusCode {
		t.Errorf("Wrong status code %d\n%s\n", response.Result().StatusCode, response.Body.String())
	}
}

func assertListEqual(t *testing.T, actual []models.User, waited []models.User) {
	if len(actual) != len(waited) {
		actualBytes, _ := json.Marshal(actual)
		waitedBytes, _ := json.Marshal(waited)
		t.Errorf("Bad response \nActual:\n%s\nWaited:\n%s\n", actualBytes, waitedBytes)
		return
	}
	for obj1 := range actual {
		match := false
		for obj2 := range waited {
			if obj1 == obj2 {
				match = true
				break
			}
		}
		if !match {
			actualBytes, _ := json.Marshal(actual)
			waitedBytes, _ := json.Marshal(waited)
			t.Errorf("Bad response \nActual:\n%s\nWaited:\n%s\n", actualBytes, waitedBytes)
		}
	}
}

func assertUserEqual(t *testing.T, actual models.User, waited models.User) {
	if actual != waited {
		actualBytes, _ := json.Marshal(actual)
		waitedBytes, _ := json.Marshal(waited)
		t.Errorf("Bad response \nActual:\n%s\nWaited:\n%s\n", actualBytes, waitedBytes)
	}
}

func TestUsers(t *testing.T) {
	controllers.ConnectTestDB()
	t.Run("Test GET All users", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "", nil)
		response := httptest.NewRecorder()
		endpoints.UsersEndpoint(response, request)
		assertStatusCode(t, response, 200)
		waitedData := controllers.GetAllUsers()
		var actualData []models.User
		json.Unmarshal(response.Body.Bytes(), &actualData)
		assertListEqual(t, actualData, waitedData)
	})
	t.Run("Test GET user by id", func(t *testing.T) {
		user := controllers.GetAllUsers()[1]
		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%d", user.ID), nil)
		response := httptest.NewRecorder()
		endpoints.UsersEndpoint(response, request)
		assertStatusCode(t, response, 200)
		var actualData models.User
		json.Unmarshal(response.Body.Bytes(), &actualData)
		assertUserEqual(t, actualData, user)
	})
	t.Run("Test POST user", func(t *testing.T) {
		userData := map[string]interface{}{}
		userData["name"] = "New user"
		userData["age"] = 2
		body, _ := json.Marshal(userData)
		request, _ := http.NewRequest(http.MethodPost, "", bytes.NewBuffer(body))
		response := httptest.NewRecorder()
		endpoints.UsersEndpoint(response, request)
		assertStatusCode(t, response, 201)
		var actualData models.User
		json.Unmarshal(response.Body.Bytes(), &actualData)
		waitedData := models.User{
			ID:   actualData.ID,
			Name: "New user",
			Age:  2,
		}
		assertUserEqual(t, actualData, waitedData)
	})
	t.Run("Test PUT user", func(t *testing.T) {
		userData := map[string]interface{}{}
		userData["name"] = "Upd user"
		userData["age"] = 2
		body, _ := json.Marshal(userData)
		request, _ := http.NewRequest(http.MethodPost, "", bytes.NewBuffer(body))
		response := httptest.NewRecorder()
		endpoints.UsersEndpoint(response, request)
		var user models.User
		json.Unmarshal(response.Body.Bytes(), &user)

		userData["age"] = 22
		body, _ = json.Marshal(userData)
		request, _ = http.NewRequest(http.MethodPut, fmt.Sprintf("%d", user.ID), bytes.NewBuffer(body))
		response = httptest.NewRecorder()
		endpoints.UsersEndpoint(response, request)
		assertStatusCode(t, response, 200)
		var actualData models.User
		json.Unmarshal(response.Body.Bytes(), &actualData)
		waitedData := models.User{
			ID:   actualData.ID,
			Name: "Upd user",
			Age:  22,
		}
		assertUserEqual(t, actualData, waitedData)
	})
	t.Run("Test DELETE user", func(t *testing.T) {
		userData := map[string]interface{}{}
		userData["name"] = "Del user"
		userData["age"] = 2
		body, _ := json.Marshal(userData)
		request, _ := http.NewRequest(http.MethodPost, "", bytes.NewBuffer(body))
		response := httptest.NewRecorder()
		endpoints.UsersEndpoint(response, request)
		var user models.User
		json.Unmarshal(response.Body.Bytes(), &user)

		request, _ = http.NewRequest(http.MethodDelete, fmt.Sprintf("%d", user.ID), nil)
		response = httptest.NewRecorder()
		endpoints.UsersEndpoint(response, request)
		assertStatusCode(t, response, 204)

		request, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("%d", user.ID), nil)
		response = httptest.NewRecorder()
		endpoints.UsersEndpoint(response, request)
		assertStatusCode(t, response, 404)
	})
}
