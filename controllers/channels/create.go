package channels

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/grumblechat/server/internal/channel"
)

type typeEnvelope struct {
	Type string `json:"type" validation:"oneof:text voice, required"`
}

func createHandler(ctx echo.Context) error {
	envelope := &typeEnvelope{}

	// unmarshal channel type
	if err := ctx.Bind(envelope); err != nil {
		return ctx.JSON(http.StatusBadRequest, envelope)
	}

	// create voice
	if envelope.Type == "voice" {
		voiceChan := channel.CreateVoice()
		if err := ctx.Bind(voiceChan); err != nil {
			return ctx.JSON(http.StatusBadRequest, voiceChan)
		}
		return ctx.JSON(http.StatusCreated, voiceChan)
	}

	// create text
	if envelope.Type == "text" {
		textChan := channel.CreateText()
		if err := ctx.Bind(textChan); err != nil {
			return ctx.JSON(http.StatusBadRequest, textChan)
		}
		return ctx.JSON(http.StatusCreated, textChan)
	}

	// handle unexpected type
	return echo.NewHTTPError(http.StatusInternalServerError)
}