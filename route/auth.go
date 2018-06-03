package route

import (
	"net/http"
	"github.com/suyashkumar2409/ChitChat/config"
	"github.com/suyashkumar2409/ChitChat/data"
)

func Login(w http.ResponseWriter, r * http.Request){
	_, err := verifySession(w, r)
	if err == nil{
		http.Redirect(w, r, indexUrl, http.StatusFound)
	}
	generateHTML(w, nil, "layout", "public.navbar", "login")
}

func Logout(w http.ResponseWriter, r * http.Request){
	cookie, err := r.Cookie(cookieName)
	if err != nil{
		config.Warning(err, "Failed to get cookie")
	} else {
		sess := data.Session{Uuid:cookie.Value}
		err := sess.DeleteByUUID()
		if err != nil{
			config.Warning(err, "Failed to delete session")
		}
	}
	http.Redirect(w, r, indexUrl, http.StatusFound)
}

func Signup(w http.ResponseWriter, r * http.Request){
	_, err := verifySession(w, r)
	if err == nil{
		http.Redirect(w, r, indexUrl, http.StatusFound)
	}
	generateHTML(w, nil, "layout", "public.navbar", "signup")
}

func SignupAccount(w http.ResponseWriter, r * http.Request){
	err := r.ParseForm()
	if err != nil{
		config.Error(err, "Could not parse form")
		errorMessage(w, r, "Error encountered with Sign Up")
	}
	user := data.User{
		Name: r.PostFormValue("name"),
		Email: r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}
	if err := user.Create() ; err != nil{
		config.Error(err, "Could not create user")
		errorMessage(w, r, "Could not create this user")
	}
	http.Redirect(w, r, loginURL, http.StatusFound)
}

func Authenticate(w http.ResponseWriter, r * http.Request){
	err := r.ParseForm()
	if err != nil{
		config.Error(err, "Could not parse form")
		errorMessage(w, r, "Error encountered while logging in")
	}
	user, err := data.GetUserByEmail(r.FormValue("email"))
	if err != nil{
		config.Warning(err, "Cannot find user")
	}
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		sess, err := user.CreateSession()
		if err != nil{
			config.Error(err, "Cannot create session")
			http.Redirect(w, r, loginURL, http.StatusFound)
		}
		cookie := http.Cookie{
			Name: cookieName,
			Value: sess.Uuid,
			HttpOnly:true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, indexUrl, http.StatusFound)
	} else {
		http.Redirect(w, r, loginURL, http.StatusFound)
	}
}



