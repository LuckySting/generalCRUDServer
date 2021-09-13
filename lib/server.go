package lib

import (
	"../controllers"
	"../endpoints"
	"net/http"
)

func RunServer(port string) {
	controllers.ConnectDB()
	http.HandleFunc("/users/", endpoints.UsersEndpoint)
	http.ListenAndServe(":"+port, nil)
}
