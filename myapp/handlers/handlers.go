package handlers

import (
	"github.com/andihoerudin24/goand"
	"net/http"
)

type Handlers struct {
	App *goand.Goand
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(w, r, "home", nil)
	if err != nil {
		h.App.ErrorLog.Println("Error Rendering", err)
	}
}
