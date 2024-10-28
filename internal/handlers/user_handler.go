package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/nurmuh-alhakim18/music-catalog-api/internal/models/user"
	"github.com/nurmuh-alhakim18/music-catalog-api/pkg/utils"
)

type userService interface {
	Register(ctx context.Context, req user.UserRegisterRequest) (user.User, error)
	Login(ctx context.Context, req user.UserLoginRequest) (user.UserLoginResponse, error)
}

type UserHandler struct {
	userService userService
}

func NewUserHandler(userService userService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) HandlerRegister(w http.ResponseWriter, r *http.Request) {
	var params user.UserRegisterRequest
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "Invalid input", err)
		return
	}

	user, err := h.userService.Register(r.Context(), params)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error(), err)
		return
	}

	utils.Response(w, http.StatusCreated, user)
}

func (h *UserHandler) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	var params user.UserLoginRequest
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "Invalid input", err)
		return
	}

	loginResp, err := h.userService.Login(r.Context(), params)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error(), err)
		return
	}

	utils.Response(w, http.StatusOK, loginResp)
}
