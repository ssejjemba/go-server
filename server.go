package rest_server

import (
	"fmt"
	"net/http"
	"strings"
)

type UserStore interface {
	GetUserName(id string) string
}

type UserServer struct {
	store UserStore
}

func (u *UserServer) ServeHTTP (w http.ResponseWriter, r *http.Request){
	switch r.Method{
		case http.MethodPost:
			u.processUser(w)
		case http.MethodGet:
			u.showUser(w, r)
	}
}

func (u *UserServer) processUser(w http.ResponseWriter){
	w.WriteHeader(http.StatusAccepted)
}

func (u *UserServer) showUser(w http.ResponseWriter, r *http.Request){
	user := strings.TrimPrefix(r.URL.Path, "/users/")


	name := u.store.GetUserName(user)

	if name == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, name)
}


