package api

import (
	"encoding/json"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/nats-io/nats.go"
)

var NoData = []byte("")

type MsgType string

const (
	MsgRunScript   MsgType = "run.script"
	MsgActorCreate MsgType = "actor.create"
	MsgActorDelete MsgType = "actor.delete"
	MsgGetValue    MsgType = "get.value"
	MsgSetValue    MsgType = "set.value"
	MsgSignal      MsgType = "signal"
	MsgCall        MsgType = "call"
	MsgGetState    MsgType = "get.state"
	MsgSetState    MsgType = "set.state"
	MsgListActors  MsgType = "list.actors"
	MsgWorldCreate MsgType = "world.create"
	MsgWorldDelete MsgType = "world.delete"
	MsgWorldStatus MsgType = "world.status"
	MsgWorldListen MsgType = "world.listen"
	MsgWorldClose  MsgType = "world.close"
	MsgWorldCall   MsgType = "world.call"
	MsgWorldEvents MsgType = "world.events"
	MsgPing        MsgType = "ping"
)

type Msg struct {
	Type   MsgType         `json:"type"`
	World  string          `json:"world"`
	Actor  string          `json:"actor"`
	Member string          `json:"member"`
	Data   json.RawMessage `json:"data"`
}

func DoRequest(nc *nats.Conn, msg *Msg) (*Msg, error) {
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	subject := "sim." + string(msg.Type)
	reply, err := nc.Request(subject, data, timout)
	if err != nil {
		return nil, err
	}
	var res Msg
	if len(reply.Data) == 0 {
		return &res, nil
	}
	err = json.Unmarshal(reply.Data, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func DoPublish(nc *nats.Conn, msg *Msg) error {
	subject := "sim." + string(msg.Type)
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return nc.Publish(subject, data)
}

func OnPublish(nc *nats.Conn, mt MsgType, fn func(msg *Msg) error) (*nats.Subscription, error) {
	subject := "sim." + string(mt)
	return nc.Subscribe(subject, func(msg *nats.Msg) {
		var m Msg
		if len(msg.Data) == 0 {
			log.Error().Msg("empty message")
			return
		}
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			return
		}
		err = fn(&m)
		if err != nil {
			return
		}
	})
}

func OnRequest(nc *nats.Conn, mt MsgType, fn func(msg *Msg) (*Msg, error)) (*nats.Subscription, error) {
	subject := "sim." + string(mt)
	return nc.Subscribe(subject, func(msg *nats.Msg) {
		var m Msg
		if len(msg.Data) == 0 {
			log.Error().Msg("empty message")
			msg.Respond(NoData)
			return
		}
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			log.Error().Err(err).Msg("failed to unmarshal message")
			msg.Respond(NoData)
			return
		}
		res, err := fn(&m)
		if err != nil {
			log.Error().Err(err).Msg("failed to handle request")
			msg.Respond(NoData)
			return
		}
		if res == nil {
			msg.Respond(NoData)
			return
		}
		data, err := json.Marshal(res)
		if err != nil {
			log.Error().Err(err).Msg("failed to marshal response")
			msg.Respond(NoData)
			return
		}
		err = msg.Respond(data)
		if err != nil {
			log.Error().Err(err).Msg("failed to respond to request")
			return
		}
	})
}
