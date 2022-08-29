package rest_server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)



func TestServer(t *testing.T) {
	t.Run("Returns the user that is being requested", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "users/1", nil)

		response := httptest.NewRecorder()

		UserServer(response, request)

		got := response.Body.String()

		want := User {
			id: "01",
			firstName: "Daniel",
			lastName: "Ssejjemba",
			age: 26,
			address: "Mego Bilania Apartments, Kyebando",
			gender: "Male",
			phone: "+256706650884",
		}

		str := fmt.Sprintf("%#v", want)

		if got != str {
			t.Errorf("got %s wanted %s", got, str)
		}
	})
}

