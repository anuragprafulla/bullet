package handlers

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/anuragprafulla/bullet/internal/users"
	"github.com/anuragprafulla/bullet/pkg/errors"
)

type IUserHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type userhandler struct {
	store users.IUserStore
}

func NewUserHandler(store users.IUserStore) IUserHandler {
	return &userhandler{store: store}
}

func (h *userhandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		WriteError(w, errors.ErrorBadRequest)
		return
	}

	user, err := h.store.Get(r.Context(), &users.GetUserRequest{ID: id})

	if err != nil {
		WriteError(w, err)
		return
	}

	WriteResponse(w, &users.UserResponseWrapper{User: user})
}

func (h *userhandler) List(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	after := values.Get("after")
	name := values.Get("names")

	limit, err := IntFromString(w, values.Get("limit"))

	if err != nil {
		return
	}
	// list users
	list, err := h.store.List(r.Context(), &users.ListUserRequest{
		Limit: limit,
		After: after,
		Name:  name,
	})

	if err != nil {
		WriteError(w, err)
		return
	}

	WriteResponse(w, &users.UserResponseWrapper{Users: list})

}

func (h *userhandler) Create(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	log.Println(data)
	if err != nil {
		WriteError(w, errors.ErrorBadRequest)
		return
	}
	user := &users.User{}

	if Unmarshal(w, data, user) != nil {
		return
	}

	if err = h.store.Create(r.Context(), &users.CreateUserRequest{User: user}); err != nil {
		WriteError(w, err)
		return
	}
	WriteResponse(w, &users.UserResponseWrapper{User: user})
}

func (h *userhandler) Update(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteError(w, errors.ErrorBadRequest)
		return
	}
	req := &users.UpdateUserRequest{}

	if Unmarshal(w, data, req) != nil {
		return
	}

	// check if event exist
	if _, err := h.store.Get(r.Context(), &users.GetUserRequest{ID: req.ID}); err != nil {
		WriteError(w, err)
		return
	}

	if err = h.store.Update(r.Context(), req); err != nil {
		WriteError(w, err)
		return
	}

	WriteResponse(w, &users.UserResponseWrapper{})
}

func (h *userhandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		WriteError(w, errors.ErrorBadRequest)
		return
	}

	// check if event exist
	if _, err := h.store.Get(r.Context(), &users.GetUserRequest{ID: id}); err != nil {
		WriteError(w, err)
		return
	}

	if err := h.store.Delete(r.Context(), &users.DeleteUserRequest{ID: id}); err != nil {
		WriteError(w, err)
		return
	}

	WriteResponse(w, &users.UserResponseWrapper{})
}
