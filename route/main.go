package route

import (
	"net/http"
	"github.com/suyashkumar2409/ChitChat/data"
	"github.com/suyashkumar2409/ChitChat/config"
)

func Index(w http.ResponseWriter, r * http.Request){
	threads, err := data.GetThreads()
	if err != nil{
		config.Error(err, "Cannot get all threads")
		errorMessage(w, r, "Cannot get threads")
	}
	_, err = verifySession(w, r)
	if err != nil{
		generateHTML(w, threads, "layout", "public.navbar", "index")
	} else {
		generateHTML(w, threads, "layout", "private.navbar", "index")
	}
}

func Err(w http.ResponseWriter, r * http.Request){
	vals := r.URL.Query()
	_, err := verifySession(w, r)
	if err != nil{
		generateHTML(w, vals.Get(errorMsg), "layout", "public.navbar", "error")
	} else {
		generateHTML(w, vals.Get(errorMsg), "layout", "private.navbar", "error")
	}
}
