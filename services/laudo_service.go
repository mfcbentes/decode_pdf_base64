package services

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/mfcbentes/decode_pdf_base64/repositories"
	"golang.org/x/exp/slog"
)

func CreateLaudos() ([]string, error) {
	laudos, err := repositories.GetLaudos()
	if err != nil {
		slog.Error("Erro ao obter laudos", slog.Any("error", err))
		return nil, err
	}

	var filePaths []string
	for _, laudo := range laudos {
		// Decodificar o Base64
		pdfData, err := base64.StdEncoding.DecodeString(laudo.DsPdfSerial)
		if err != nil {
			slog.Error("Erro ao decodificar Base64", slog.Any("error", err))
			return nil, fmt.Errorf("erro ao decodificar Base64: %v", err)
		}

		// Salvar o PDF em um arquivo
		filePath := fmt.Sprintf("/app/output/%d.pdf", laudo.NrAcessoDicom)
		file, err := os.Create(filePath)
		if err != nil {
			slog.Error("Erro ao criar arquivo PDF", slog.String("filePath", filePath), slog.Any("error", err))
			return nil, fmt.Errorf("erro ao criar arquivo PDF: %v", err)
		}
		defer file.Close()

		_, err = file.Write(pdfData)
		if err != nil {
			slog.Error("Erro ao escrever no arquivo PDF", slog.String("filePath", filePath), slog.Any("error", err))
			return nil, fmt.Errorf("erro ao escrever no arquivo PDF: %v", err)
		}

		slog.Info("Arquivo PDF criado com sucesso", slog.String("filePath", filePath))
		filePaths = append(filePaths, filePath)
	}

	return filePaths, nil
}
