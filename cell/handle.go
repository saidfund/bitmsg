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
	log.Infof("%s recv HAND | %+v\n", self.Name, ev.Message())
	sack := &SackMsg{
		Code: "200",
		Info: "ok",
	}
	ev.Session().Send(sack)

}

func (self *Handler) HandleSack(ev cellnet.Event) {
	log.Infof("%s recv SACK | %+v\n", self.Name, ev.Message())
	ev.Session().Close()

}
