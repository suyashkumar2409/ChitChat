package route

import (
	"net/http"
	"github.com/suyashkumar2409/ChitChat/config"
	"github.com/suyashkumar2409/ChitChat/data"
	"fmt"
)

func NewThread(w http.ResponseWriter, r * http.Request){
	_, err := verifySession(w, r)
	if err != nil{
		http.Redirect(w, r, loginURL, http.StatusFound)
	} else {
		generateHTML(w, nil, "layout", "private.navbar", "new.thread")
	}
}


func CreateThread(w http.ResponseWriter, r * http.Request){
	sess, err := verifySession(w, r)
	if err != nil{
		http.Redirect(w, r, loginURL, http.StatusFound)
	} else {
		err = r.ParseForm()
		if err != nil{
			config.Error(err, "Cannot parse form")
			http.Redirect(w, r, indexUrl, http.StatusFound)
		}
		user, err := sess.GetUser()
		if err != nil{
			config.Error(err, "Cannot get user from session")
			http.Redirect(w, r, indexUrl, http.StatusFound)
		}
		topic := r.PostFormValue("topic")
		if _, err := user.CreateThread(topic) ; err != nil{
			config.Error(err, "Cannot create thread")
			http.Redirect(w, r, indexUrl, http.StatusFound)
		}
		http.Redirect(w, r, indexUrl, http.StatusFound)
	}
}

func PostThread(w http.ResponseWriter, r * http.Request){
	sess, err := verifySession(w, r)
	if err != nil{
		http.Redirect(w, r, loginURL, http.StatusFound)
	} else {
		err = r.ParseForm()
		if err != nil{
			config.Error(err, "Cannot parse form")
			http.Redirect(w, r, indexUrl, http.StatusFound)
		}
		user, err := sess.GetUser()
		if err != nil{
			config.Error(err, "Cannot get user from session")
			http.Redirect(w, r, indexUrl, http.StatusFound)
		}
		body := r.PostFormValue("body")
		uuid := r.PostFormValue("uuid")
		thread, err := data.GetThreadByUUID(uuid)
		if err != nil{
			config.Error(err, "Cannot read thread")
			http.Redirect(w, r, indexUrl, http.StatusFound)
		}
		if _, err := user.CreatePost(thread, body) ; err != nil {
			config.Error(err, "Cannot create post")
			http.Redirect(w, r, indexUrl, http.StatusFound)
		}
		url := fmt.Sprintf("/thread/read?id=%s", uuid)
		http.Redirect(w, r, url, http.StatusFound)
	}
}

func ReadThread(w http.ResponseWriter, r * http.Request){
	uuid := r.URL.Query().Get("id")
	thread, err := data.GetThreadByUUID(uuid)
	if err != nil{
		config.Warning(err, "Could not find thread")
		errorMessage(w, r, "Could not find thread")
	} else {
		_, err := verifySession(w, r)
		if err != nil{
			generateHTML(w, &thread, "layout", "public.navbar", "public.thread")
		} else {
			generateHTML(w, &thread, "layout", "private.navbar", "private.thread")
		}
	}
}