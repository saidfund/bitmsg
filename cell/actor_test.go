package cell

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/golog"
	"testing"
	"time"
)

func TestCell(t *testing.T) {
	golog.SetLevelByString("tcpproc", "info")

	address := "127.0.0.1:17701"

	done := make(chan struct{})
	queue := cellnet.NewEventQueue()
	queue.StartLoop()
	server := &Handler{}
	server.Role = "Server"
	go server.Acceptor(queue, address)

	time.Sleep(2 * time.Second)

	client := &Handler{}
	client.Role = "Client"
	go client.Connect(queue, address)
	<-done
}
