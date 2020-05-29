package interfaces

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// IMysqlDB IMysqlDB
type IMysqlDB interface {
	GetDB() *sql.DB
}
