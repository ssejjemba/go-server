package rest_server

import (
	"fmt"
	"net/http"
)

type User struct {
	id string
	firstName string
	lastName string
	age int
	address string
	gender string
	phone string
}

func UserServer(w http.ResponseWriter, r *http.Request){
	 dummy_user := User {
			"01",
			"Daniel",
			"Ssejjemba",
			26,
			"Mego Bilania Apartments, Kyebando",
			"Male",
			"+256706650884",
		}
	fmt.Fprint(w, dummy_user)
}