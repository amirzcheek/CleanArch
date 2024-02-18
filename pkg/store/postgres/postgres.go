package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type DBInfo struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func (db *DBInfo) toDSN() string {
	var args []string

	if len(db.Host) > 0 {
		args = append(args, fmt.Sprintf("host=%s", db.Host))
	}

	if len(db.Port) > 0 {
		args = append(args, fmt.Sprintf("port=%s", db.Port))
	}

	if len(db.User) > 0 {
		args = append(args, fmt.Sprintf("user=%s", db.User))
	}

	if len(db.Password) > 0 {
		args = append(args, fmt.Sprintf("password=%s", db.Password))
	}

	if len(db.DBName) > 0 {
		args = append(args, fmt.Sprintf("dbname=%s", db.DBName))
	}

	args = append(args, "sslmode=disable")

	return strings.Join(args, " ")
}

func ConnectDB(info *DBInfo) (*sql.DB, error) {
	connStr := info.toDSN()

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
