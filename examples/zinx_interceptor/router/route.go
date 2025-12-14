package router

import (
	"github.com/kruily/zinx/ziface"
	"github.com/kruily/zinx/zlog"
	"github.com/kruily/zinx/znet"
)

type HelloRouter struct {
	znet.BaseRouter
}

func (hr *HelloRouter) Handle(request ziface.IRequest) {
	zlog.Ins().InfoF(string(request.GetData()))
}
