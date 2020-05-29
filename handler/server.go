package handler

import (
	"github.com/mikudos/mikudos_golang_app/interfaces"
)

// Server 事件驱动服务间流程控制方法，提供基本的数据库操作方法
type Hstruct struct {
	Server interfaces.IServer
}

func (h Hstruct) ServerDecorater(interfaces.IServer) {}
