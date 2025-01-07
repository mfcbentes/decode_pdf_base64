package tasks

import (
	"time"

	"github.com/mfcbentes/decode_pdf_base64/services"
	"golang.org/x/exp/slog"
)

func GenerateLaudosPeriodically() {
	// Função para gerar laudos
	generateLaudos := func() {
		_, err := services.CreateLaudos()
		if err != nil {
			slog.Error("Erro ao criar PDFs", slog.Any("error", err))
		} else {
			slog.Info("Laudos gerados com sucesso")
		}
	}

	generateLaudos()

	// Configura um ticker para executar a geração de laudos a cada 10 minutos
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			generateLaudos()
		}
	}()
}
