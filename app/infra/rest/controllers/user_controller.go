package controllers

import (
	"context"
	"encoding/json"
	"github.com/inaohiro-sandbox/wire-tutorial4/app/domain"
	"log"
	"net/http"
	"strings"
)

type UserController struct {
	usecase domain.UserService
}

func NewUserController(usecase domain.UserService) *UserController {
	return &UserController{usecase: usecase}
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// get user name from request params
	//name := r.Body.Read()

	log.Printf(strings.TrimPrefix(r.URL.Path, "users"))

	name := "john"
	created, err := c.usecase.Create(ctx, name)
	if err != nil {
		log.Printf(err.Error())
		// todo: return error as json
	}

	_json, err := json.Marshal(created)
	if err != nil {
		log.Printf(err.Error())
		// todo: return errors as json
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(_json)
}
