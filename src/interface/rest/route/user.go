package route

import (
	"net/http"

	handlers "ex_service/src/interface/rest/handlers/user"

	"github.com/go-chi/chi/v5"
)

func UserRouter(h handlers.UserHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Post("/create_user", h.Register)
	r.Post("/login", h.Login)

	return r
}
