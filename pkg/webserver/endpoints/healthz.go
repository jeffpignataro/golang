package endpoints

import (
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"
)

type HealthzHandler struct {
	log *zap.Logger
}

func NewHealthzHandler(log *zap.Logger) *HealthzHandler {
	return &HealthzHandler{log: log}
}

func (*HealthzHandler) Pattern() string {
	return "/healthz"
}

func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Error("Failed to read request", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if _, err := fmt.Fprint(w, "I am alive!!"); err != nil {
		h.log.Error("Failed to write response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
