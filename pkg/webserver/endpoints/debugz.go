package endpoints

import (
	"io"
	"net/http"

	"go.uber.org/zap"
)

type DebugzHandler struct {
	log *zap.Logger
}

func NewDebugzHandler(log *zap.Logger) *DebugzHandler {
	return &DebugzHandler{log: log}
}

func (*DebugzHandler) Pattern() string {
	return "/debugz"
}

func (h *DebugzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Debug endpoint initiated")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Debug endpoint initiated.")
}
