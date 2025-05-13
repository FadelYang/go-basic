package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type Project struct {
	Stack    string `json:"stack"`
	Platform string `json:"platform"`
}

func main() {
	fullpath := "https://www.fadelanumah.com/project?stack=react&platform=website"

	e, _ := url.Parse(fullpath)

	fmt.Println("scheme:", e.Scheme)
	fmt.Println("host:", e.Host)
	fmt.Println("path:", e.Path)
	fmt.Println("query:", e.Query())

	query := e.Query()

	fmt.Println("stack:", query["stack"][0])
	fmt.Println("platform:", query["platform"][0])

	DecodeJSONToObject()
	DecodeJSONToArray()
	DecodeJSONToMapStringInterface()
	DecodeJSONToInterface()
	EncodeObjectToJSON()
}

func DecodeJSONToObject() {
	var project Project

	jsonString := `{
		"stack": "Laravel",
		"platform": "website"
	}`

	jsonData := []byte(jsonString)
	err := json.Unmarshal(jsonData, &project)
	if err != nil {
		fmt.Println("failed to get json data", err)
	}

	fmt.Println("Become object:")
	fmt.Println(project.Stack)
	fmt.Println(project.Platform)
}

func DecodeJSONToArray() {
	var project []Project

	jsonArray := `[
		{
			"stack": "laravel",
			"platform": "website"
		},
		{
			"stack": "bootstrap",
			"platform": "website"
		},
		{
			"stack": "flutter",
			"platform": "Android"
		}
	]`

	jsonData := []byte(jsonArray)
	err := json.Unmarshal(jsonData, &project)
	if err != nil {
		fmt.Println("failed to get json data", err)
	}

	fmt.Println("Become array:")
	fmt.Println(project)
}

func DecodeJSONToMapStringInterface() {
	var project map[string]interface{}

	jsonArray := `{
			"stack": "laravel",
			"platform": "website"
		}
	`

	jsonData := []byte(jsonArray)
	err := json.Unmarshal(jsonData, &project)
	if err != nil {
		fmt.Println("failed to get json data", err)
	}

	fmt.Println("Become map of interface:")
	fmt.Println(project)
}

func DecodeJSONToInterface() {
	var project interface{}

	jsonArray := `{
			"stack": "laravel",
			"platform": "website"
		}
	`

	jsonData := []byte(jsonArray)
	err := json.Unmarshal(jsonData, &project)
	if err != nil {
		fmt.Println("failed to get json data", err)
	}

	fmt.Println("Become interface:")
	fmt.Println(project)

	decodedProjectdata := project.(map[string]interface{})
	fmt.Println("After decoded")
	fmt.Println(project)
	fmt.Println(decodedProjectdata["stack"])
}

func EncodeObjectToJSON() {
	data := Project{"react", "website"}
	res, _ := json.Marshal(data)

	fmt.Println(string(res))
}
