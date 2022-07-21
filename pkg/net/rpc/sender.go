package rpc

import (
	"encoding/json"
	"net/url"

	"github.com/gorilla/websocket"
)

type RpcSender struct {
	conn   *websocket.Conn
	writer RpcMessageHandler
}

func NewRpcSender(writer RpcMessageHandler) *RpcSender {
	return &RpcSender{
		writer: writer,
	}
}

func (s *RpcSender) Dial(addr string) error {
	u := url.URL{Scheme: "ws", Host: addr, Path: "/ws/"}
	log.Infof("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}
	s.conn = c
	return nil
}

func (s *RpcSender) Close() {
	s.conn.Close()
}

func (s *RpcSender) SendMessages(emitter chan RpcMessage) {
	for message := range emitter {
		log.Infof("send  %v\n", message)
		err := s.conn.WriteJSON(message)
		if err != nil {
			log.Warnf("failed to send message: %s", err)
		}
	}
}

func (s *RpcSender) ReadPump() {
	defer func() {
		log.Debug("readPump: close")
		s.conn.Close()
	}()
	for {
		_, data, err := s.conn.ReadMessage()
		if err != nil {
			log.Debugf("readPump: %v", err)
			return
		}
		log.Debugf("received message: %s", data)
		var msg RpcMessage
		err = json.Unmarshal(data, &msg)
		if err != nil {
			log.Debugf("error decoding rpc message: %v", err)
			return
		}
		log.Debugf("decoded message: %+v", msg)
		err = s.writer.HandleMessage(msg)
		if err != nil {
			log.Debugf("error writing message: %v", err)
			return
		}
	}
}
