package rest_server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)



func TestServer(t *testing.T) {
	t.Run("Returns the user with id 1", func(t *testing.T) {
		request := newGetUserRequest("1")

		response := httptest.NewRecorder()

		UserServer(response, request)

		got := response.Body.String()

		assertResponseBody(t, got, "Daniel")
	})

	t.Run("Returns the user with id 2", func(t *testing.T) {
		request := newGetUserRequest("2")

		response := httptest.NewRecorder()

		UserServer(response, request)

		got := response.Body.String()

		assertResponseBody(t, got, "Jordan")
	})
}

func newGetUserRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", name), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

