package auth

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gerry-sheva/bts-todo-list/pkg/apierror"
	"github.com/gerry-sheva/bts-todo-list/pkg/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthHandler struct {
	logger *slog.Logger
	dbpool *pgxpool.Pool
}

// Creates new AuthHandler with logger and dbpool as the dependencies
func New(logger *slog.Logger, dbpool *pgxpool.Pool) *AuthHandler {
	return &AuthHandler{
		logger,
		dbpool,
	}
}

// Registers new user
func (h *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var i AuthInput
	if err := util.ReadJSON(w, r, &i); err != nil {
		apierror.GlobalErrorHandler.BadRequestResponse(w, r, err)
		return
	}

	h.logger.Info("Registering new user", "username", i.Username)

	v := util.NewValidator()

	if i.validate(v); !v.Valid() {
		apierror.GlobalErrorHandler.FailedValidationResponse(w, r, v.Errors)
		return
	}

	exp := time.Now().Add(time.Hour * 24 * 7)

	jwt, err := register(r.Context(), h.dbpool, &i)
	if err != nil {
		apierror.GlobalErrorHandler.ServerErrorResponse(w, r, err)
		return
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    jwt,
		Expires:  exp,
		Secure:   true,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	err = util.WriteJSON(w, http.StatusOK, util.Envelope{"jwt": jwt}, nil)
	if err != nil {
		apierror.GlobalErrorHandler.ServerErrorResponse(w, r, err)
	}

	h.logger.Info("Successfully registered new user", "username", i.Username)
}

func (h *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var i AuthInput
	if err := util.ReadJSON(w, r, &i); err != nil {
		apierror.GlobalErrorHandler.BadRequestResponse(w, r, err)
		return
	}

	h.logger.Info("Logging in new user", "username", i.Username)

	v := util.NewValidator()

	if i.validate(v); !v.Valid() {
		apierror.GlobalErrorHandler.FailedValidationResponse(w, r, v.Errors)
		return
	}

	exp := time.Now().Add(time.Hour * 24 * 7)
	jwt, err := login(r.Context(), h.dbpool, &i)
	if err != nil {
		apierror.GlobalErrorHandler.ServerErrorResponse(w, r, err)
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    jwt,
		Expires:  exp,
		Secure:   true,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	err = util.WriteJSON(w, http.StatusOK, util.Envelope{"jwt": jwt}, nil)
	if err != nil {
		apierror.GlobalErrorHandler.ServerErrorResponse(w, r, err)
	}
}
