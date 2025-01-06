package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/godror/godror"
)

func ConnectDB() (*sql.DB, error) {
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
