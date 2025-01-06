package views

import (
	"fmt"
	"net/http"
)

func RenderPDF(w http.ResponseWriter, r *http.Request, filePath string, sequence int) {
	// Definir cabeçalhos para exibir o PDF no navegador
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=%d.pdf", sequence))
	http.ServeFile(w, r, filePath)
}
