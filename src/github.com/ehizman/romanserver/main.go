package main

import (
	"fmt"
	"html"
	"net/http"
	"github.com/ehizman"
	"strconv"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		urlPathElements := strings.Split(request.URL.Path, "/")
		if urlPathElements[1] == "roman_number" {
			number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
			if number == 0 || number >= 10 {
				writer.WriteHeader(http.StatusNotFound)
				writer.Write([]byte("404- Not Found"))
			} else {
				fmt.Fprintf(writer, "%q",
					html.EscapeString(romanNumerals.Numerals[number]))
			}
		} else {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte("400 - Bad Request"))
		}
	})

	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
