package rpc

import (
	"apigear/pkg/log"
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
		log.Debugf("send message %+v\n", message)
		s.conn.WriteJSON(message)
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
			log.Debugf("error decoding rpc mesage: %v", err)
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
