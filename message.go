package main

type PubMessage struct {
	topic string
	msg   string
}

func (m PubMessage) getTopic() string {
	return m.topic
}

func (m PubMessage) getMsg() string {
	return m.msg
}

type SubMessageChan struct {
	topic string
	msgCh chan *PubMessage
}

func (s *SubMessageChan) getTopic() string {
	return s.topic
}

func (s *SubMessageChan) getMsgCh() chan *PubMessage {
	return s.msgCh
}
