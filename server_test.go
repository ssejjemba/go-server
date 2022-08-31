package rest_server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubUserStore struct {
	names map[string] string
}

func (s *StubUserStore) GetUserName(id string) string {
	return s.names[id]
}

func TestServer(t *testing.T) {
	store := StubUserStore {
		map[string]string{
			"1": "Daniel",
			"2": "Jordan",
		},
	}
	server := &UserServer{ &store }

	t.Run("Returns the user with id 1", func(t *testing.T) {
		request := newGetUserRequest("1")

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, got, "Daniel")
	})

	t.Run("Returns the user with id 2", func(t *testing.T) {
		request := newGetUserRequest("2")

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, got, "Jordan")
	})

	t.Run("returns 404 on missing users", func(t *testing.T) {
		request := newGetUserRequest("5")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreUsers(t *testing.T) {
	store := StubUserStore{
		map[string]string{},
	}
	server := &UserServer{&store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/users/1", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
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

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

