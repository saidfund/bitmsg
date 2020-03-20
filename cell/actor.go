package cell

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	_ "github.com/davyxu/cellnet/peer/tcp"
	"github.com/davyxu/cellnet/proc"
	_ "github.com/davyxu/cellnet/proc/tcp"
	"net"
)

const peerAddress = "127.0.0.1:17701"

type Handler struct {
	Name string
	Role string
	Link string
}

func (self *Handler) Acceptor(queue cellnet.EventQueue, addr string) {
	done := make(chan struct{})
	//queue := cellnet.NewEventQueue()
	self.Role = "Server"
	acceptor := peer.NewGenericPeer("tcp.Acceptor", "server", addr, queue)
	proc.BindProcessorHandler(acceptor, "tcp.ltv", self.HandleEvent)
	acceptor.Start()
	<-done
}

func (self *Handler) Connect(queue cellnet.EventQueue, workerAddr string) {
	// 等待服务器返回数据
	self.Role = "Client"
	done := make(chan struct{})
	//queue := cellnet.NewEventQueue()
	connector := peer.NewGenericPeer("tcp.Connector", "client", workerAddr, queue)
	proc.BindProcessorHandler(connector, "tcp.ltv", self.HandleEvent)
	connector.Start()
	//queue.StartLoop()
	// 等待客户端收到消息
	<-done
}

func (self *Handler) SetLink(ev cellnet.Event) string {
	if self.Link != "" {
		return self.Link
	}
	conn, ok := ev.Session().Raw().(net.Conn)
	if !ok {
		self.Link = ""
		return self.Link
	}
	raddr := conn.RemoteAddr().String()
	laddr := conn.LocalAddr().String()
	if self.Role == "Server" {
		self.Link = raddr + "<-" + laddr
	} else {
		self.Link = laddr + "->" + raddr
	}
	return self.Link
}

func (self *Handler) OnAccepted(ev cellnet.Event) {
	self.SetLink(ev)
	log.Infoln(self.Role, "server accepted ", self.Link)
}

func (self *Handler) OnConnected(ev cellnet.Event) {
	self.SetLink(ev)
	log.Infoln(self.Role, "client connected ", self.Link)
	self.ShakeHand(ev)
}

func (self *Handler) OnClosed(ev cellnet.Event) {
	log.Infoln(self.Role, "Closed :", self.Link)
}

func (self *Handler) HandleEvent(ev cellnet.Event) {
	switch ev.Message().(type) {
	case *cellnet.SessionAccepted:
		self.OnAccepted(ev)
	case *cellnet.SessionConnected: // 已经连接上
		self.OnConnected(ev)

	case *cellnet.SessionClosed:
		self.OnClosed(ev)
	}
}
