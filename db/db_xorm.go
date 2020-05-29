package db

import (
	"log"

	"github.com/go-xorm/xorm"
)

// Database Database
type Database struct {
	engine *xorm.Engine
}

// GetEngine GetEngine
func (db *Database) GetEngine() *xorm.Engine {
	return db.engine
}

// DB DB
var DB Database

func init() {
	x, err := xorm.NewEngine("mysql", "root:111111@/sys?charset=utf8")
	if err != nil {
		log.Fatalf("Database connect error")
	}
	DB = Database{x}
}
