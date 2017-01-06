package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func Server(w http.ResponseWriter, r *http.Request) {
	min := 1000
	minS := r.URL.Query().Get("min")
	if minS != "" {
		min, _ = strconv.Atoi(minS)
	}

	i := 0

	cont := true
	for cont {
		i++
		cont = false
		for y := 2; y < i; y++ {
			if i%y == 0 {
				cont = true
			}
		}

		if i <= min {
			cont = true
		}
	}

	io.WriteString(w, strconv.Itoa(i)+"\n")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "1234"
	}
	http.HandleFunc("/", Server)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
