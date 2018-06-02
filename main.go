package main

import (
	"net/http"
	"blog/route"
	"fmt"
	"blog/config"
	"time"
)

func createMuxAndRegisterHandlers() (*http.ServeMux) {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.GetStatic()))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", route.Index)
	mux.HandleFunc("/err", route.Err)


	mux.HandleFunc("/login", route.Login)
	mux.HandleFunc("/logout", route.Logout)
	mux.HandleFunc("/signup", route.Signup)
	mux.HandleFunc("/signup_account", route.SignupAccount)
	mux.HandleFunc("/authenticate", route.Authenticate)

	mux.HandleFunc("/thread/new", route.NewThread)
	mux.HandleFunc("/thread/create", route.CreateThread)
	mux.HandleFunc("/thread/post", route.PostThread)
	mux.HandleFunc("/thread/read", route.ReadThread)

	return mux
}

func main() {
	fmt.Println("Chitchat", config.GetVersion(), "started at", config.GetAddress())

	mux := createMuxAndRegisterHandlers()

	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
		ReadTimeout:time.Duration(config.GetReadTimeout() * int64(time.Second)),
		WriteTimeout:time.Duration(config.GetWriteTimeout() * int64(time.Second)),
		MaxHeaderBytes: 1<<20,
	}
	server.ListenAndServe()
}