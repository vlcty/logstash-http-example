package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func manglePlainText(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	requestBodyContent, _ := io.ReadAll(request.Body)

	response.WriteHeader(http.StatusOK)
	fmt.Fprintln(response, "Received:", string(requestBodyContent))
}

func mangleJSON(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	type Input struct {
		Number string `json:"number"`
		Animal string `json:"animal"`
	}

	input := &Input{}
	json.NewDecoder(request.Body).Decode(input)

	type Result struct {
		DoubledNumber  int    `json:"number_doubled"`
		RepeatedAnimal string `json:"animal_repeated"`
	}

	result := &Result{}
	inputNumber, _ := strconv.Atoi(input.Number)
	result.DoubledNumber = inputNumber * 2
	result.RepeatedAnimal = strings.Repeat(input.Animal, 2)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(result)
}

func main() {
	http.HandleFunc("/plain", manglePlainText)
	http.HandleFunc("/json", mangleJSON)

	http.ListenAndServe(":8080", nil)
}
