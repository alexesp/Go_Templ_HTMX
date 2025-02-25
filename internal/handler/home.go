package handler

import (
	"net/http"

	"github.com/alexesp/Go_Templ_HTMX.git/internal/view/home"
)

type homeHandler struct{}

func (h homeHandler) handleIndex(w http.ResponseWriter, r *http.Request) error {
	// if _, err := io.WriteString(w, "Hola Go!"); err != nil {
	// 	return err
	//}
	//return nil

	return home.Index().Render(r.Context(), w)
}
