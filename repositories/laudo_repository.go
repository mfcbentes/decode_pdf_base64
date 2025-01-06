package repositories

import (
	"fmt"

	"github.com/mfcbentes/decode_pdf_base64/config"
	"github.com/mfcbentes/decode_pdf_base64/models"
)

func GetLaudo(sequence int) (*models.Laudo, error) {
	config.LoadEnv()

	db, err := config.ConnectDB()
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar no banco: %v", err)
	}
	defer db.Close()

	var dsPdfSerial string

	err = db.QueryRow("SELECT DS_PDF_SERIAL FROM laudo_paciente_pdf_serial WHERE nr_acesso_dicom = :1", sequence).Scan(&dsPdfSerial)
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar: %v", err)
	}

	return &models.Laudo{Sequence: sequence, DsPdfSerial: dsPdfSerial}, nil
}
