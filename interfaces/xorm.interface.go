package interfaces

import "github.com/go-xorm/xorm"

// IXorm Xorm interface
type IXorm interface {
	GetEngine() *xorm.Engine
}
