package rest

import (
	"github.com/inaohiro-sandbox/wire-tutorial4/app/infra/rest/controllers"
	"net/http"
)

type Router http.Handler

// Router provider returns http.Handler, which is used in server.go
// Also, router requires controllers, so that it requires to pass controller as arguments
func NewRouter(userController *controllers.UserController) Router {
	mux := http.NewServeMux()
	mux.HandleFunc("/users/", userController.Create)

	return mux
}
