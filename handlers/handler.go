package handlers

import "github.com/mphox-phoxdev/kobold-generator/kobold"

// Handler is a helper-struct to pass config info relevant to the mux-requests
type Handler struct {
	KoboldDB kobold.Database
}
