package user

import (
	"encoding/json"
	"log"
	"net/http"

	dto "ex_service/src/app/dto/user"
	usecases "ex_service/src/app/usecases/user"
	common_error "ex_service/src/infra/errors"
	"ex_service/src/interface/rest/response"

	"github.com/go-chi/chi/v5"
	"github.com/lib/pq"
)

type UserHandlerInterface interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	LoginSocialMedia(w http.ResponseWriter, r *http.Request)
	LoginSocialMediaCallback(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	response response.IResponseClient
	usecase  usecases.UserUCInterface
}

func NewUserandler(r response.IResponseClient, h usecases.UserUCInterface) UserHandlerInterface {
	return &userHandler{
		response: r,
		usecase:  h,
	}
}

func (h *userHandler) LoginSocialMedia(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	data, err := h.usecase.LoginSocialMedia(provider)
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	h.response.JSON(w, "Login Social Media", data, nil)
}
func (h *userHandler) Register(w http.ResponseWriter, r *http.Request) {

	postDTO := dto.RegisterReqDTO{}
	err := json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil {
		log.Println(err)

		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}
	err = postDTO.Validate()
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	data, err := h.usecase.Register(&postDTO)
	if err != nil {
		log.Println(err)
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			h.response.HttpError(w, common_error.NewError(common_error.USER_ALREADY_EXIST, err))
			return
		}
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_CREATE_DATA, err))
		return
	}
	w.WriteHeader(http.StatusCreated)
	h.response.JSON(
		w,
		"Successful Register New User",
		data,
		nil,
	)
}

func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) {

	postDTO := dto.LoginReqDTO{}
	err := json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}
	err = postDTO.Validate()
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	data, err := h.usecase.Login(&postDTO)
	if err != nil {
		log.Println(err)

		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Login",
		data,
		nil,
	)
}

func (h *userHandler) LoginSocialMediaCallback(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	if provider == "google" {
		h.LoginWithGoogleCallback(w, r)
		return
	}

	h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, nil))
}

func (h *userHandler) LoginWithGoogleCallback(w http.ResponseWriter, r *http.Request) {
	// get code from query params
	code := r.URL.Query().Get("code")

	// get user info from google
	userInfo, err := h.usecase.ExchangeCodeGoogle(r.Context(), code)
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Login with Google",
		userInfo,
		nil,
	)
}
