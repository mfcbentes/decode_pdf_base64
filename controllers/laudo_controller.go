package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/mfcbentes/decode_pdf_base64/views"
	"golang.org/x/exp/slog"
)

func HandleLaudo(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/laudo/")
	if !strings.HasSuffix(path, ".pdf") {
		slog.Error("Formato de URL inválido", slog.String("path", path))
		http.Error(w, "Formato de URL inválido", http.StatusBadRequest)
		return
	}

	sequenceStr := strings.TrimSuffix(path, ".pdf")
	sequence, err := strconv.Atoi(sequenceStr)
	if err != nil {
		slog.Error("Parâmetro 'sequence' inválido", slog.String("sequenceStr", sequenceStr), slog.Any("error", err))
		http.Error(w, "Parâmetro 'sequence' inválido", http.StatusBadRequest)
		return
	}

	filePath := fmt.Sprintf("/app/output/%d.pdf", sequence)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		slog.Error("Arquivo PDF não encontrado", slog.String("filePath", filePath))
		http.Error(w, "Arquivo PDF não encontrado", http.StatusNotFound)
		return
	}

	views.RenderPDF(w, r, filePath, sequence)
}

func HandleStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("API on-line"))
}
