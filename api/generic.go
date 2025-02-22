package api

import (
	"net/http"
)

type Server struct {
	Debug bool
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s.Debug {
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	switch r.URL.Path {
	case "/account/info":
		s.handleAccountInfo(w, r)
	case "/account/register":
		s.handleAccountRegister(w, r)
	case "/account/login":
		s.handleAccountLogin(w, r)
	case "/account/logout":
		s.handleAccountLogout(w, r)

	case "/game/playercount":
		s.handlePlayerCountGet(w)

	case "/savedata/get":
		s.handleSavedataGet(w, r)
	case "/savedata/update":
		s.handleSavedataUpdate(w, r)
	case "/savedata/delete":
		s.handleSavedataDelete(w, r)
	case "/savedata/clear":
		s.handleSavedataClear(w, r)

	case "/daily/seed":
		s.handleSeed(w)
	case "/daily/rankings":
		s.handleRankings(w, r)
	case "/daily/rankingpagecount":
		s.handleRankingPageCount(w, r)
	}
}

// auth

type GenericAuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GenericAuthResponse struct {
	Token string `json:"token"`
}
