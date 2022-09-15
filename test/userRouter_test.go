package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go-app/router"
)

func TestIUserGetRouter(t *testing.T) {
	type User struct {
		UserId string `json:"userId"`
	}
	user := User{
		UserId: "123",
	}
	expectedBody, _ := json.Marshal(user)
	router := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user/"+user.UserId, nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expectedBody), w.Body.String())
}