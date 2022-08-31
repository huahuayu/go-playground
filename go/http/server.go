package http

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validate = validator.New()

func hello(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
	default:
		ResponseErr(w, ErrMethodNotAllowed)
		return
	}
	if !req.URL.Query().Has("name") {
		ResponseErr(w, ErrInvalidParam)
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
	switch req.Method {
	case http.MethodPost:
		user := User{}
		err := json.NewDecoder(req.Body).Decode(&user)
		if err != nil {
			ResponseErr(w, ErrInvalidParam)
			return
		}
		err = validate.Struct(user)
		if err != nil {
			ResponseErr(w, ErrInvalidParam, err.Error())
			return
		}
		ResponseOK(w, user)
	default:
		ResponseErr(w, ErrMethodNotAllowed)
		return
	}
}

func serve() {
	InitMsgLanguage(ENGLISH)
	// get
	http.HandleFunc("/hello", hello)
	// post
	http.HandleFunc("/user", user)
	err := http.ListenAndServe(":6000", nil)
	if err != nil {
		panic(err)
	}
}
