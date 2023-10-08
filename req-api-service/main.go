package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func main() {
	for {
		waterValue := rand.Intn(100) + 1
		windValue := rand.Intn(100) + 1

		data := map[string]interface{} {
			"water": waterValue,
			"wind": windValue,
		}

		requestJson, err := json.Marshal(&data)
		client := http.Client{}

		if err != nil {
			log.Fatalln(err)
		}

		request, err := http.NewRequest("PUT", "http://localhost:8080/weather", bytes.NewBuffer(requestJson))

		request.Header.Set("Content-Type", "application/json")

		if err != nil {
			log.Fatalln(err)
		}

		response, err := client.Do(request)

		if err != nil {
			log.Fatalln(err)
		}

		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)

		if err != nil {
			log.Fatalln(err)
		}

		if err != nil {
			log.Fatalln(err)
		}

		var temp map[string]interface{}

		err = json.Unmarshal([]byte(body), &temp)

		if err != nil {
			log.Fatalln(err)
		}

		weather, err := json.Marshal(temp["weather"])

		if err != nil {
			log.Fatalln(err)
		}

		result := fmt.Sprintf(`
			%+v
			status water: %s
			status wind: %s
		`, string(weather), temp["water_status"], temp["wind_status"])

		cleanedStringResult := strings.Replace(result, "\t", "", -1)

		log.Println(cleanedStringResult)

		time.Sleep(time.Second * 15)
	}
}