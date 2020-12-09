package handlers

import (
	"github.com/DmitryKuzmenec/ImgPrettify/services"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

// Handler
type Handler struct {
	svc *services.Svc
}

// NewImgHandler
func NewImgHandler(svc *services.Svc) *Handler {
	return &Handler{
		svc: svc,
	}
}

// Pretty
func (h *Handler) Pretty(ctx echo.Context) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		log.Errorf("[handler] ctx.FormFile: %s", err)
		return err
	}
	newFile, err := h.svc.Pretty(file)
	if err != nil {
		log.Errorf("[handler] svc.Pretty: %s", err)
		return err
	}
	return ctx.File(newFile.Name())
}
