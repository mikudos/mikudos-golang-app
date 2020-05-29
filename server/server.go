package server

import (
	"github.com/mikudos/mikudos_golang_app/interfaces"
	pb "github.com/mikudos/mikudos_golang_app/proto/greeter"
	cron "github.com/robfig/cron/v3"
	"github.com/yueguanyu/go_debugger"
)

// Handler Server instance
var Handler Server

// Server Server
type Server struct {
	pb.GreeterServiceServer
	Cron        *cron.Cron
	CronJobs    map[cron.EntryID]string
	OneTimeJobs map[cron.EntryID]string
	SceneName   string
	Database    interfaces.IDB
	Handlers    interfaces.IHandler
}

var debug = go_debugger.Debug("mikudos:server")
