package repositories

import (
	"fmt"

	"github.com/mfcbentes/decode_pdf_base64/config"
	"github.com/mfcbentes/decode_pdf_base64/models"
)

func GetLaudos() ([]models.Laudo, error) {
	config.LoadEnv()

	db, err := config.ConnectDB()
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar no banco: %v", err)
	}
	defer db.Close()

	query := `
    SELECT
        initcap(obter_nome_pf(ap.cd_pessoa_fisica))           AS nm_paciente,
        TRIM(obter_desc_proc_interno(pp.nr_seq_proc_interno)) AS ds_procedimento,
        ap.cd_pessoa_fisica                                   AS protocolo,
        pp.nr_seq_interno                                     AS senha,
        obter_compl_pf(ap.cd_pessoa_fisica, 1, 'DDDCEL')
        || obter_compl_pf(ap.cd_pessoa_fisica, 1, 'CEL')      AS nr_telefone,
        ge.nr_prescricao,
        ap.nr_atendimento,
        lpps.nr_acesso_dicom,
        lpps.DS_PDF_SERIAL
    FROM
        eis_gestao_exames_v  ge,
        prescr_procedimento  pp,
        prescr_medica        pm,
        atendimento_paciente ap,
        laudo_paciente_pdf_serial lpps
    WHERE
        trunc(lpps.dt_atualizacao) = trunc(sysdate)
        AND ge.nr_prescricao = pp.nr_prescricao
        AND ge.nr_seq_proced = pp.nr_sequencia
        AND pm.nr_prescricao = pp.nr_prescricao
        AND pm.nr_atendimento = ap.nr_atendimento
        AND pp.nr_seq_interno = lpps.nr_acesso_dicom
    ORDER BY
        pp.dt_baixa`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar a consulta: %v", err)
	}
	defer rows.Close()

	var laudos []models.Laudo
	for rows.Next() {
		var laudo models.Laudo
		err := rows.Scan(
			&laudo.NmPaciente,
			&laudo.DsProcedimento,
			&laudo.Protocolo,
			&laudo.Senha,
			&laudo.NrTelefone,
			&laudo.NrPrescricao,
			&laudo.NrAtendimento,
			&laudo.NrAcessoDicom,
			&laudo.DsPdfSerial,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear os resultados: %v", err)
		}
		laudos = append(laudos, laudo)
	}

	return laudos, nil
}
