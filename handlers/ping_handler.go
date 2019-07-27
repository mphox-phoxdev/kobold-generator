package handlers

import (
	"net/http"

	"github.com/mphox-phoxdev/kobold-generator/core"
)

func (h *Handler) PingHandler(w http.ResponseWriter, r *http.Request) {
	core.LogIncomingRequest(r)
	core.WriteJSONResponse(w, "pong")
}
