package handler

import (
	"io"
	"net/http"
)

type homeHandler struct{}

func (h homeHandler) handleIndex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hola Go!")
}
