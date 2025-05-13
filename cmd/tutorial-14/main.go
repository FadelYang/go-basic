package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Project struct {
	Id       string `json:"id"`
	Stack    string `json:"stack"`
	Platform string `json:"platform"`
}

var projectData = []Project{
	{
		Id:       "prj001",
		Stack:    "react, tailwind",
		Platform: "website",
	},
	{
		Id:       "prj002",
		Stack:    "laravel, vue",
		Platform: "website",
	},
	{
		Id:       "prj003",
		Stack:    "flutter",
		Platform: "smartphone",
	},
}

func main() {
	http.HandleFunc("/projects", GetProject)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println("Can not start the service, something went wrong:", err)
	} else {
		fmt.Println("Success running in port 5000")
	}
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	projectPlatform := r.URL.Query().Get("platform")

	filteredProjectData := projectData

	var temp []Project
	if r.Method == "GET" {
		if projectPlatform != "" {
			for _, project := range projectData {
				if project.Platform == projectPlatform {
					temp = append(temp, project)
				}
			}
			filteredProjectData = temp
		}

		res, err := json.Marshal(filteredProjectData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(res)
		return
	}
}
