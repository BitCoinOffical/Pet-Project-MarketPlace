package main

import (
	"access_manager-service/config"
	"access_manager-service/internal/auth"
	"access_manager-service/internal/db"
	"access_manager-service/internal/handlers"
	rediscash "access_manager-service/internal/redis"
	"access_manager-service/pkg/code"
	"access_manager-service/pkg/search"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	database, err := db.UsersDB(cfg)
	if err != nil {
		log.Fatal("failed conected in database: ", err)
		return
	}
	rdb, err := rediscash.RedisCashInit()
	if err != nil {
		log.Fatal("failed connect redis: ", err)
		return
	}
	reg := auth.NewRegistation(database)
	mailcode := code.NewSaveCodeInRedis(rdb)
	searcher := search.NewSearcher(database)
	h := handlers.NewHandler(searcher, mailcode, reg, cfg)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /register", h.UserRegisterHandler)
	// аунтификация пользователя по коду
	mux.HandleFunc("POST /auth/byemail", h.AuthUserByEmailHandler)
	mux.HandleFunc("POST /auth/byemail/verify", h.VerifyUserHandler)
	// аунтификация пользователя по паролю
	mux.HandleFunc("POST /auth/bypass", h.AuthUserByPassHandler)

	log.Fatal(http.ListenAndServe(cfg.ApiPort, mux))
}
