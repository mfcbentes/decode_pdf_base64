package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/godror/godror"
	"golang.org/x/exp/slog"
)

func ConnectDB() (*sql.DB, error) {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	connectString := os.Getenv("CONNECT_STRING")

	if username == "" || password == "" || connectString == "" {
		slog.Error("Variáveis de ambiente DB_USER, DB_PASSWORD ou CONNECT_STRING não definidas")
		return nil, fmt.Errorf("variáveis de ambiente DB_USER, DB_PASSWORD ou CONNECT_STRING não definidas")
	}

	dsn := fmt.Sprintf("%s/%s@%s", username, password, connectString)
	connParams, err := godror.ParseDSN(dsn)
	if err != nil {
		slog.Error("Erro ao analisar DSN", slog.String("dsn", dsn), slog.Any("error", err))
		return nil, err
	}

	connParams.Timezone = time.FixedZone("BRT", -3*60*60)

	connector := godror.NewConnector(connParams)
	db := sql.OpenDB(connector)
	slog.Info("Conexão com o banco de dados estabelecida")
	return db, nil
}
