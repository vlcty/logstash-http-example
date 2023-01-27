package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/dostuff", func(response http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()

		requestBodyContent, _ := io.ReadAll(request.Body)

		response.WriteHeader(http.StatusOK)
		fmt.Fprintln(response, "Received:", string(requestBodyContent))
	})

	http.ListenAndServe(":8080", nil)
}
