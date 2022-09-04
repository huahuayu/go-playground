package http

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validate = validator.New()

func hello(w http.ResponseWriter, req *http.Request) {
	locale := req.Header.Get("Locale")
	if req.Method != http.MethodGet {
		ResponseErr(w, locale, ErrMethodNotAllowed)
		return
	}
	if !req.URL.Query().Has("name") {
		ResponseErr(w, locale, ErrInvalidParam)
		return
	}
	name := req.URL.Query().Get("name")
	ResponseOK(w, "hello "+name)
}

func user(w http.ResponseWriter, req *http.Request) {
	type User struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Gender   string `json:"gender,omitempty"`
	}
	locale := req.Header.Get("Locale")
	if req.Method != http.MethodPost {
		ResponseErr(w, locale, ErrMethodNotAllowed)
		return
	}
	user := User{}
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		ResponseErr(w, locale, ErrInvalidParam)
		return
	}
	err = validate.Struct(user)
	if err != nil {
		ResponseErr(w, locale, ErrInvalidParam, err.Error())
		return
	}
	ResponseOK(w, user)
}

func Serve() {
	// Get
	http.HandleFunc("/hello", hello)
	// Post
	http.HandleFunc("/user", user)
	// File server
	// to Serve a directory on disk (/tmp) under an alternate URL
	// path (/file/), use StripPrefix to modify the request
	// URL's path before the FileServer sees it:
	http.Handle("/file/", http.StripPrefix("/file/", http.FileServer(http.Dir("/tmp"))))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
