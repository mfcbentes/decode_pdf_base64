package services

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/mfcbentes/decode_pdf_base64/repositories"
)

func CreatePDF(sequence int) (string, error) {
	laudo, err := repositories.GetLaudo(sequence)
	if err != nil {
		return "", err
	}

	// Decodificar o Base64
	pdfData, err := base64.StdEncoding.DecodeString(laudo.DsPdfSerial)
	if err != nil {
		return "", fmt.Errorf("erro ao decodificar Base64: %v", err)
	}

	// Salvar o PDF em um arquivo
	filePath := fmt.Sprintf("/app/output/%d.pdf", sequence)
	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("erro ao criar arquivo PDF: %v", err)
	}
	defer file.Close()

	_, err = file.Write(pdfData)
	if err != nil {
		return "", fmt.Errorf("erro ao escrever no arquivo PDF: %v", err)
	}

	log.Printf("Arquivo PDF criado com sucesso: %s", filePath)
	return filePath, nil
}
