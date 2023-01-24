package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"packform-test/config"

	"github.com/goccy/go-json"
	"github.com/joho/godotenv"
)

var token string
var base_uri string

type Response struct {
	Data string
}

func GetToken() []byte {
	creds := map[string]interface{}{"identity": config.Config("USERNAME"), "password": config.Config("PASSWORD")}
	jsonStr, _ := json.Marshal(creds)
	response, err := http.Post(base_uri+"/auth/login", "application/json", bytes.NewBuffer(jsonStr))

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return content
}

func uploadLargeFile(uri string, filePath string, chunkSize int, params map[string]string) []byte {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	file_stats, _ := file.Stat()
	defer file.Close()

	//use pipe to pass request
	rd, wr := io.Pipe()
	defer rd.Close()

	// Write file to buffer in separate go routine
	go func() {
		defer wr.Close()

		//write file
		buf := make([]byte, chunkSize)
		for {
			n, err := file.Read(buf)
			if err != nil {
				break
			}
			_, _ = wr.Write(buf[:n])
		}
	}()

	req, _ := http.NewRequest("POST", base_uri+uri, rd)
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "binary/octet-stream")
	req.ContentLength = file_stats.Size()

	// Send req using http Client
	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return content
}

func ImportCompanies() {
	res := string(uploadLargeFile("/company/import", "/home/zaid/Desktop/test_data/Test task - Postgres - customer_companies.csv", 50000, nil))
	println(res)
}

func ImportCustomers() {
	res := string(uploadLargeFile("/customer/import", "/home/zaid/Desktop/test_data/Test task - Postgres - customers.csv", 50000, nil))
	println(res)
}

func ImportOrders() {
	res := string(uploadLargeFile("/order/import", "/home/zaid/Desktop/test_data/Test task - Postgres - orders.csv", 50000, nil))
	println(res)
}

func ImportOrderItems() {
	res := string(uploadLargeFile("/order-item/import", "/home/zaid/Desktop/test_data/Test task - Postgres - order_items.csv", 50000, nil))
	println(res)
}

func ImportDeliveryItems() {
	res := string(uploadLargeFile("/delivery/import", "/home/zaid/Desktop/test_data/Test task - Postgres - deliveries.csv", 50000, nil))
	println(res)
}

func GenerateReports() {
	req, _ := http.NewRequest("GET", base_uri+"/report/refresh?start_time=2020-01-01T00:00:00.000Z&end_time=2024-01-01T00:00:00.000Z", nil)
	req.Header.Add("Authorization", "Bearer "+token)

	// Send req using http Client
	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	content, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	println(string(content))
}

func main() {
	// Load .env file
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}

	base_uri = "http://localhost:3000/api"
	var response Response
	response_bytes := GetToken()
	json.Unmarshal(response_bytes, &response)
	token = response.Data
	fmt.Println(token)
	ImportCompanies()
	ImportCustomers()
	ImportOrders()
	ImportOrderItems()
	ImportDeliveryItems()
	GenerateReports()
}
