package main

import (
	"fmt"
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
	fmt.Fprintln(w, fizzbuzz(n))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("listening on :8080")
	http.ListenAndServe(":8080", nil)
}
