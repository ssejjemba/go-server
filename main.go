package rest_server

import (
	"log"
	"net/http"
)

type InMemoryUserStore struct{}

func (i *InMemoryUserStore) GetUserName(name string) string {
	return "Live"
}

func main() {
	server := &UserServer{ &InMemoryUserStore{} }
	log.Fatal(http.ListenAndServe(":5000", server))
}