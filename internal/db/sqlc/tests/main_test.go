package sqlc__test

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"tg-backend/internal/db/sqlc"

	_ "github.com/lib/pq"
)

var testQueries *sqlc.Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open("postgres", "host=localhost dbname=yoba sslmode=disable")
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}
	testQueries = sqlc.New(conn)
	os.Exit(m.Run())
}
