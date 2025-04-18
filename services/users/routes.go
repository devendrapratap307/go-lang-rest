package users

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/restapi-go/auth"
	"github.com/restapi-go/types"
	"github.com/restapi-go/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")

	// router.HandleFunc("/users", h.CreateUser).Methods("POST")
	// router.HandleFunc("/users/{id}", h.GetUser).Methods("GET")
	// router.HandleFunc("/users/{id}", h.UpdateUser).Methods("PUT")
	// router.HandleFunc("/users/{id}", h.DeleteUser).Methods("DELETE")
	// router.HandleFunc("/users", h.ListUsers).Methods("GET")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginUserPayload
	// parse JSON
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request payload: %v", err))
		return
	}
	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("payload validation error : %v", errors))
		return
	}

	// check if user exists
	user, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error checking user: %v", err))
		return
	}
	if user == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user doesnt exists: %v", user))
		return
	}
	// Check password

	errcom := auth.ComparePassword(user.Password, payload.Password)
	if errcom != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("wrong password: %v", errcom))
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"login": "success"})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterPayload
	// parse JSON
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request payload: %v", err))
		return
	}
	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("payload validation error : %v", errors))
		return
	}

	// check if user already exists
	user, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error checking user: %v", err))
		return
	}
	if user != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user already exists: %v", user))
		return
	}
	// hash password

	hashedPassword, err := auth.CreateHashedPasssword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error hashing password: %v", err))
		return
	}
	// register user
	uErr := h.store.CreateUser(&types.User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: hashedPassword,
	})
	if uErr != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error creating user: %v", uErr))
		return
	}

	// return success response
	// utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "user created successfully"})
	utils.WriteJSON(w, http.StatusOK, map[string]bool{"register": true})

}
