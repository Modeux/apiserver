package databases

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"runtime"
	"time"
)

func MysqlConn() (*sqlx.DB, error) {
	conn := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOTS"),
		os.Getenv("DB_DATABASE"),
	)
	db, err := sqlx.Open("mysql", conn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	numCPU := runtime.NumCPU()
	db.SetMaxOpenConns(numCPU)
	db.SetMaxIdleConns(numCPU)
	db.SetConnMaxLifetime(3 * time.Minute)
	return db, nil
}
