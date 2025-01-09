package main

import (
	"net/http"
	"os"
	"time"

	"github.com/mfcbentes/decode_pdf_base64/controllers"
	"github.com/mfcbentes/decode_pdf_base64/tasks"
	"golang.org/x/exp/slog"
)

func setupLogging() {
	location, err := time.LoadLocation("America/Santarem")
	if err != nil {
		slog.Error("Erro ao carregar localização", slog.Any("error", err))
		os.Exit(1)
	}
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Value = slog.StringValue(time.Now().In(location).Format("2006/01/02 15:04:05"))
			}
			return a
		},
	})))
}

func main() {
	setupLogging()

	// Executa a geração de laudos imediatamente na inicialização
	tasks.GenerateLaudos()

	// Configura um ticker para executar a geração de laudos a cada 10 minutos
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			tasks.GenerateLaudos()
		}
	}()

	http.HandleFunc("/", controllers.HandleStatus)
	http.HandleFunc("/laudo/", controllers.HandleLaudo)
	slog.Info("Servidor iniciado na porta 8080")
	slog.Error("Erro no servidor HTTP", slog.Any("error", http.ListenAndServe(":8080", nil)))
}
