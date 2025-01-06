package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/mfcbentes/decode_pdf_base64/views"
)

func HandleLaudo(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/laudo/")
	if !strings.HasSuffix(path, ".pdf") {
		http.Error(w, "Formato de URL inválido", http.StatusBadRequest)
		return
	}

	sequenceStr := strings.TrimSuffix(path, ".pdf")
	sequence, err := strconv.Atoi(sequenceStr)
	if err != nil {
		http.Error(w, "Parâmetro 'sequence' inválido", http.StatusBadRequest)
		return
	}

	filePath := fmt.Sprintf("/app/output/%d.pdf", sequence)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "Arquivo PDF não encontrado", http.StatusNotFound)
		return
	}

	views.RenderPDF(w, r, filePath, sequence)
}
