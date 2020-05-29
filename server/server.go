package server

import (
	"github.com/mikudos/mikudos_golang_app/interfaces"
	pb "github.com/mikudos/mikudos_golang_app/proto/greeter"
	cron "github.com/robfig/cron/v3"
	"github.com/yueguanyu/go_debugger"
)

// Server Server instance
var Server Struct

// Struct ServerStruct
type Struct struct {
	pb.GreeterServiceServer
	Cron        *cron.Cron
	CronJobs    map[cron.EntryID]string
	OneTimeJobs map[cron.EntryID]string
	SceneName   string
	MongoDB     *interfaces.IMongoDB
	XormDB      *interfaces.IXorm
	MysqlDB     *interfaces.IMysqlDB
	Handlers    *interfaces.IHandler
}

var debug = go_debugger.Debug("mikudos:server")

// SetDB SetDB
func (s *Struct) SetDB(db interface{}) error {
	md, ok := db.(*interfaces.IMongoDB)
	if ok {
		s.MongoDB = md
		return nil
	}
	xd, ok := db.(*interfaces.IXorm)
	if ok {
		s.XormDB = xd
		return nil
	}
	mysqld, ok := db.(*interfaces.IMysqlDB)
	if ok {
		s.MysqlDB = mysqld
		return nil
	}
	return ERR_UNEXPECTED_DB_TYPE
}

// GetDB GetDB
func (s *Struct) GetDB(dbType string) (interface{}, error) {
	switch dbType {
	case "mongodb":
		return s.MongoDB, nil
	case "xorm":
		return s.XormDB, nil
	case "mysql":
		return s.MysqlDB, nil
	}
	return nil, ERR_UNEXPECTED_DB_TYPE
}

// Configure Configure
func (s *Struct) Configure(h *interfaces.IHandler) {
	s.Handlers = h
	(*h).Configure(s)
}
