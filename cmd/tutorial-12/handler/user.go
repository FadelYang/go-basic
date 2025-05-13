package handler

import (
	"fmt"
	"go-basic-alex-mux/cmd/tutorial-12/model"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("cmd/tutorial-12/html/index.html")
	if err != nil {
		fmt.Println(err)
	}

	user := model.User{
		Name: "Fadela",
		Job:  "Web Developer",
	}

	template.Execute(w, user)
}
