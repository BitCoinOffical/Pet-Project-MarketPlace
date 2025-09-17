package handlers

import (
	"access_manager-service/config"
	"access_manager-service/internal/auth"
	"access_manager-service/internal/storage/dto"
	"access_manager-service/pkg/code"
	passwordhash "access_manager-service/pkg/hash"
	"access_manager-service/pkg/httphelper"
	"access_manager-service/pkg/search"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	searcher search.UserSearcher
	mailcode code.CodeSaver
	reg      auth.Registrator
	cfg      *config.Config
}

func NewHandler(searcher search.UserSearcher, mailcode code.CodeSaver, reg auth.Registrator, cfg *config.Config) *Handler {
	return &Handler{
		searcher: searcher,
		mailcode: mailcode,
		reg:      reg,
		cfg:      cfg,
	}
}

func (h *Handler) UserRegisterHandler(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()

	var regDTO *dto.Register

	if err := json.NewDecoder(r.Body).Decode(&regDTO); err != nil {
		http.Error(w, "failed decode in struct", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(regDTO); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if err := h.reg.UserRegister(r.Context(), regDTO); err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			http.Error(w, "email or username already exists", http.StatusConflict)
			return
		}
		http.Error(w, "failed user register", http.StatusInternalServerError)
		return
	}

	if err := httphelper.RespondJSON(w, http.StatusCreated, map[string]string{"message": "register"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// отправка кода на почту
func (h *Handler) AuthUserByEmailHandler(w http.ResponseWriter, r *http.Request) {
	var authemail dto.LoginWithEmailCode
	if err := json.NewDecoder(r.Body).Decode(&authemail); err != nil {
		http.Error(w, "failed decode in struct", http.StatusBadRequest)
		return
	}

	res, err := h.searcher.FindUserByEmail(r.Context(), authemail.Email)
	if err != nil || res.ID == 0 {
		http.Error(w, "failed find user", http.StatusBadRequest)
		return
	}

	VerificationCode, err := code.GenerateVerificationCode()
	if err != nil {
		http.Error(w, "failed get code", http.StatusInternalServerError)
		return
	}

	h.mailcode.SaveCodeRedis(r.Context(), authemail.Email, VerificationCode)

	if err := auth.GetMassageByEmail(h.cfg, VerificationCode, authemail.Email); err != nil {
		http.Error(w, "failed get code on mail", http.StatusInternalServerError)
		return
	}

	httphelper.RespondJSON(w, http.StatusOK, map[string]string{"status": "код отправлен"})

}

// получение кода от юзера
func (h *Handler) VerifyUserHandler(w http.ResponseWriter, r *http.Request) {
	var Code dto.LoginWithEmailCode

	if err := json.NewDecoder(r.Body).Decode(&Code); err != nil {
		http.Error(w, "failed decode in struct", http.StatusBadRequest)
		return
	}

	key := fmt.Sprintf("LoginCode:%s", Code.Email)
	if err := h.mailcode.CheckCodeRedis(r.Context(), key, Code.Code); err != nil {
		http.Error(w, "code not succes", http.StatusForbidden)
		return
	}

	user, err := h.searcher.FindUserByEmail(r.Context(), Code.Email)
	if err != nil {
		http.Error(w, "find not succes", http.StatusForbidden)
		return
	}

	if err := h.mailcode.CreateSessionToken(w, r.Context(), user.ID); err != nil {
		http.Error(w, "failed created session token", http.StatusBadRequest)
		return
	}

	httphelper.RespondJSON(w, http.StatusCreated, map[string]string{"status": "token created"})
}

// вход с помощью пароля
func (h *Handler) AuthUserByPassHandler(w http.ResponseWriter, r *http.Request) {
	var authusername dto.LoginWithPassword
	username, RealUserPass, err := h.searcher.FindUserByUsername(r.Context(), authusername.UserName)
	if err != nil || username == "" {
		http.Error(w, "failed find user", http.StatusBadRequest)
		return
	}
	if passwordhash.CheckPasswordHash(authusername.Password, RealUserPass) {

	}
}
