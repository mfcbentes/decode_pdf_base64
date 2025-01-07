package services

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"time"

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
		// Verificar se o arquivo já existe
		filePath := fmt.Sprintf("/app/output/%d.pdf", laudo.NrAcessoDicom)
		if _, err := os.Stat(filePath); err == nil {
			//slog.Info("Arquivo PDF já existe", slog.String("filePath", filePath))
			continue
		}

		// Decodificar o Base64
		pdfData, err := base64.StdEncoding.DecodeString(laudo.DsPdfSerial)
		if err != nil {
			slog.Error("Erro ao decodificar Base64", slog.Any("error", err))
			return nil, fmt.Errorf("erro ao decodificar Base64: %v", err)
		}

		// Salvar o PDF em um arquivo
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

func RemoveOldPDFs() error {
	outputDir := "/app/output"
	files, err := os.ReadDir(outputDir)
	if err != nil {
		return fmt.Errorf("erro ao ler o diretório: %v", err)
	}

	now := time.Now()
	for _, file := range files {
		filePath := filepath.Join(outputDir, file.Name())
		info, err := file.Info()
		if err != nil {
			slog.Error("Erro ao obter informações do arquivo", slog.String("filePath", filePath), slog.Any("error", err))
			continue
		}

		if now.Sub(info.ModTime()).Hours() > 24 {
			err := os.Remove(filePath)
			if err != nil {
				slog.Error("Erro ao remover arquivo", slog.String("filePath", filePath), slog.Any("error", err))
			} else {
				slog.Info("Arquivo removido com sucesso", slog.String("filePath", filePath))
			}
		}
	}

	return nil
}
