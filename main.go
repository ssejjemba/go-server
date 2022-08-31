package rest_server

import (
	"log"
	"net/http"
)

type InMemoryUserStore struct{
	names map[string] string
}

func (i *InMemoryUserStore) GetUserName(id string) string {
	return i.names[id]
}

func (i *InMemoryUserStore) AddUser(user UserRequest) {
	i.names[user.id] = user.name
}

func main() {
	server := &UserServer{ &InMemoryUserStore{} }
	log.Fatal(http.ListenAndServe(":5000", server))
}