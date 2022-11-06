package endpoints

import "go.uber.org/zap"

type Handler struct {
	route string
	log   *zap.Logger
}

func NewHandler(log *zap.Logger) *Handler {
	return &Handler{log: log}
}

func (h *Handler) Pattern() string {
	return h.route
}
