package rest_server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type UserRequest struct {
	id string
	name string
}

type UserStore interface {
	GetUserName(id string) string
	AddUser(user UserRequest)
}

type UserServer struct {
	store UserStore
}

func (u *UserServer) ServeHTTP (w http.ResponseWriter, r *http.Request){
	switch r.Method{
		case http.MethodPost:
			u.processUser(w, r)
		case http.MethodGet:
			u.showUser(w, r)
	}
}

func (u *UserServer) processUser(w http.ResponseWriter, r *http.Request){
	var userReq UserRequest

	// Try to Decode the request body into the struct
	err := json.NewDecoder(r.Body).Decode(&userReq)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u.store.AddUser(userReq)

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "User: %+v", userReq)
}

func (u *UserServer) showUser(w http.ResponseWriter, r *http.Request){
	user := strings.TrimPrefix(r.URL.Path, "/users/")


	name := u.store.GetUserName(user)

	if name == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, name)
}


