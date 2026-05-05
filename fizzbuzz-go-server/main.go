package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func fizzbuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return fmt.Sprintf("%d", n)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.URL.Query().Get("n"))
	if err != nil || n < 1 {
		http.Error(w, "provide ?n=<positive integer>", http.StatusBadRequest)
		return
	}
	if _, err := fmt.Fprintln(w, fizzbuzz(n)); err != nil {
		log.Printf("write error: %v", err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("fizzbuzz server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
