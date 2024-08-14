package databases

import (
	"github.com/jmoiron/sqlx"
)

const Mysql = "mysql"

type DBInterface interface {
	GetConn(name string) *sqlx.DB
}

type DB struct {
	Mysql *sqlx.DB
}

func (d *DB) GetConn(name string) *sqlx.DB {
	if name == Mysql {
		return d.Mysql
	}
	return nil
}

func NewDB() (DBInterface, error) {
	my, err := MysqlConn()
	if err != nil {
		return nil, err
	}
	db := DB{Mysql: my}
	return &db, nil
}
