package utils

import (
	"database/sql"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

func NewDB(host string, port int, user, pass, dbName string) (*sql.DB, error) {
	strReplacer := strings.NewReplacer("{user}", user, "{pass}", pass, "{host}", host, "{port}", strconv.Itoa(port), "{db}", dbName)

	dbSource := strReplacer.Replace("postgresql://{user}:{pass}@{host}:{port}/{db}?sslmode=disable")
	conn, err := sql.Open("postgres", dbSource)

	if err != nil {
		return nil, err
	}

	return conn, conn.Ping()
}
