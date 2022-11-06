package endpoints

import (
	"net/http"
	"os"

	"go.uber.org/zap"
)

type QuitQuitQuitHandler struct {
	log *zap.Logger
}

func NewQuitQuitQuitHandler(log *zap.Logger) *QuitQuitQuitHandler {
	return &QuitQuitQuitHandler{log: log}
}

func (*QuitQuitQuitHandler) Pattern() string {
	return "/quitquitquit"
}

func (h *QuitQuitQuitHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Error("Quiting application")
	os.Exit(1)
}
