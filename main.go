package main

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/godror/godror"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar .env: %v", err)
	}
}

func connectDB() (*sql.DB, error) {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	connectString := os.Getenv("CONNECT_STRING")

	if username == "" || password == "" || connectString == "" {
		return nil, fmt.Errorf("variáveis de ambiente DB_USER, DB_PASSWORD ou CONNECT_STRING não definidas")
	}

	dsn := fmt.Sprintf("%s/%s@%s", username, password, connectString)
	connParams, err := godror.ParseDSN(dsn)
	if err != nil {
		return nil, err
	}

	connParams.Timezone = time.FixedZone("BRT", -3*60*60)

	connector := godror.NewConnector(connParams)
	db := sql.OpenDB(connector)
	return db, nil
}

func createPDF(sequence int) (string, error) {
	loadEnv()

	db, err := connectDB()
	if err != nil {
		return "", fmt.Errorf("erro ao conectar no banco: %v", err)
	}
	defer db.Close()

	var dsPdfSerial string

	err = db.QueryRow("SELECT DS_PDF_SERIAL FROM laudo_paciente_pdf_serial WHERE nr_acesso_dicom = :1", sequence).Scan(&dsPdfSerial)
	if err != nil {
		return "", fmt.Errorf("erro ao consultar: %v", err)
	}

	// Decodificar o Base64
	pdfData, err := base64.StdEncoding.DecodeString(dsPdfSerial)
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

func handleLaudo(w http.ResponseWriter, r *http.Request) {
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

	// Definir cabeçalhos para exibir o PDF no navegador
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=%d.pdf", sequence))
	http.ServeFile(w, r, filePath)
}

func main() {
	// Criação do arquivo PDF na inicialização do programa
	sequence := 26554
	_, err := createPDF(sequence)
	if err != nil {
		log.Fatalf("Erro ao criar PDF: %v", err)
	}

	// Configuração do servidor HTTP
	http.HandleFunc("/laudo/", handleLaudo)
	log.Println("Servidor iniciado na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
