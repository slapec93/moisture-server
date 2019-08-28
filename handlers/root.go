package handlers

import (
	"fmt"
	"net/http"
)

// Root ...
func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Println("called")
}
