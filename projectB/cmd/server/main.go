package main

import (
	"fmt"
	"net/http"

	"github.com/hermanbanken/example-bazel-monorepo/projectB/pkg/lib"
)

func main() {
	fmt.Println("Hello from projectB")
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from projectB, a calculation: %d", lib.Calc(1, 41))
	}))
}
