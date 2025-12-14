package c_router

import (
	"github.com/kruily/zinx/ziface"
	"github.com/kruily/zinx/zlog"
	"github.com/kruily/zinx/znet"
)

type HelloRouter struct {
	znet.BaseRouter
}

// HelloZinxRouter Handle
func (this *HelloRouter) Handle(request ziface.IRequest) {
	zlog.Debug("Call HelloZinxRouter Handle")
	zlog.Debug("recv from server : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

}
