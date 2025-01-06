package main

import (
	"log"
	"net/http"

	"github.com/mfcbentes/decode_pdf_base64/controllers"
	"github.com/mfcbentes/decode_pdf_base64/services"
)

func main() {
	// Criação do arquivo PDF na inicialização do programa
	sequence := 26554
	_, err := services.CreatePDF(sequence)
	if err != nil {
		log.Fatalf("Erro ao criar PDF: %v", err)
	}

	// Configuração do servidor HTTP
	http.HandleFunc("/laudo/", controllers.HandleLaudo)
	log.Println("Servidor iniciado na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
