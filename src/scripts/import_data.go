package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/goccy/go-json"
)

var token string
var base_uri string

type Response struct {
	Data string
}

func GetToken() []byte {
	var jsonStr = []byte(`{"identity":"mzaidannas","password":"Mzaidannas007"}`)
	response, err := http.Post(base_uri+"/auth/login", "application/json", bytes.NewBuffer(jsonStr))

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return content
}

func SendPostRequest(api string, filename string) []byte {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file_stats, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	req, _ := http.NewRequest("POST", base_uri+api, file)
	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "binary/octet-stream")
	req.Header.Add("Content-Length", fmt.Sprint(file_stats.Size()))
	req.Header.Add("Accept", "application/json")

	// Send req using http Client
	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return content
}

func ImportCompanies() {
	res := string(SendPostRequest("/company/import", "/home/zaid/Desktop/test_data/Test task - Postgres - customer_companies.csv"))
	println(res)
}

func main() {
	base_uri = "http://localhost:3000/api"
	var response Response
	response_bytes := GetToken()
	json.Unmarshal(response_bytes, &response)
	token = response.Data
	fmt.Println(token)
	ImportCompanies()
}
