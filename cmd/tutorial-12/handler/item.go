package handler

import (
	"fmt"
	"net/http"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Kamu berada di halaman Item")
}
