package rest_server

import (
	"fmt"
	"net/http"
	"strings"
)


func UserServer(w http.ResponseWriter, r *http.Request){
	user := strings.TrimPrefix(r.URL.Path, "/users/")

	fmt.Fprint(w, GetUserName(user))
	
}

func GetUserName(id string) string {
	if(id == "1"){
		return "Daniel"
	}

	if(id == "2"){
		return "Jordan"
	}

	return ""

}