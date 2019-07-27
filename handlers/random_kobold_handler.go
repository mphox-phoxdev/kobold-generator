package handlers

import (
	"net/http"

	"github.com/mphox-phoxdev/kobold-generator/core"
	"github.com/mphox-phoxdev/kobold-generator/kobold"
)

func (h *Handler) RandomKoboldHandler(w http.ResponseWriter, r *http.Request) {
	core.LogIncomingRequest(r)
	core.WriteJSONResponse(w, kobold.GenerateKobold())
}
