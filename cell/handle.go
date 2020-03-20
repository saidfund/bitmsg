package cell

import "github.com/davyxu/cellnet"

func (self *Handler) ShakeHand(ev cellnet.Event) {
	hand := &HandMsg{
		Addr: "abc",
		Info: "info",
	}
	log.Infoln("send handshake")
	ev.Session().Send(hand)
}

func (self *Handler) HandleHand(ev cellnet.Event) {

}
