package message

type PubMessage struct {
	topic string
	msg   interface{}
}

func NewPubMessage(topic string, msg interface{}) *PubMessage {
	return &PubMessage{
		topic: topic,
		msg:   msg,
	}
}

func (m PubMessage) GetTopic() string {
	return m.topic
}

func (m PubMessage) GetMsg() interface{} {
	return m.msg
}

type SubMessageChan struct {
	topic string
	msgCh chan *PubMessage
}

func NewSubMessageChan(topic string) *SubMessageChan {
	return &SubMessageChan{
		topic: topic,
		msgCh: make(chan *PubMessage),
	}

}

func (s *SubMessageChan) GetTopic() string {
	return s.topic
}

func (s *SubMessageChan) GetMsgCh() chan *PubMessage {
	return s.msgCh
}
