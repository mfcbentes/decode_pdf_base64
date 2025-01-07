package tasks

import (
	"github.com/mfcbentes/decode_pdf_base64/services"
	"golang.org/x/exp/slog"
)

func GenerateLaudos() {
	_, err := services.CreateLaudos()
	if err != nil {
		slog.Error("Erro ao criar PDFs", slog.Any("error", err))
	} else {
		slog.Info("Laudos gerados com sucesso")
	}
}
