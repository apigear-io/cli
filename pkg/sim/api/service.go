package api

import (
	"encoding/json"
	"fmt"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/sim/model"
	"github.com/nats-io/nats.go"
)

// Service represents a NATS-based service that provides access to the simulation system
type Service struct {
	nc      *nats.Conn
	js      nats.JetStreamContext
	manager SimulationManager
	subs    []*nats.Subscription
	unsubs  []func()
}

// NewService creates a new NATS service with the given manager
func NewService(url string, manager SimulationManager) (*Service, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}

	js, err := nc.JetStream()
	if err != nil {
		return nil, fmt.Errorf("failed to create JetStream context: %w", err)
	}

	s := &Service{
		nc:      nc,
		js:      js,
		manager: manager,
	}

	if err := s.setup(); err != nil {
		return nil, err
	}

	return s, nil
}

// Start initializes all the NATS services
func (s *Service) setup() error {
	sub, err := OnRequest(s.nc, MsgRunScript, s.runScript)
	if err != nil {
		return fmt.Errorf("failed to subscribe to run script requests: %w", err)
	}
	s.subs = append(s.subs, sub)

	sub, err = OnRequest(s.nc, MsgWorldCreate, s.createWorld)
	if err != nil {
		return fmt.Errorf("failed to subscribe to create world requests: %w", err)
	}
	s.subs = append(s.subs, sub)

	sub, err = OnRequest(s.nc, MsgWorldDelete, s.deleteWorld)
	if err != nil {
		return fmt.Errorf("failed to subscribe to delete world requests: %w", err)
	}
	s.subs = append(s.subs, sub)

	sub, err = OnRequest(s.nc, MsgWorldStatus, s.worldStatus)
	if err != nil {
		return fmt.Errorf("failed to subscribe to world status requests: %w", err)
	}
	s.subs = append(s.subs, sub)

	sub, err = OnRequest(s.nc, MsgGetState, s.getActorState)
	if err != nil {
		return fmt.Errorf("failed to subscribe to get actor state requests: %w", err)
	}
	s.subs = append(s.subs, sub)
	sub, err = OnRequest(s.nc, MsgGetValue, s.getActorValue)
	if err != nil {
		return fmt.Errorf("failed to subscribe to get actor value requests: %w", err)
	}
	s.subs = append(s.subs, sub)
	sub, err = OnRequest(s.nc, MsgSetValue, s.setActorValue)
	if err != nil {
		return fmt.Errorf("failed to subscribe to set actor value requests: %w", err)
	}
	s.subs = append(s.subs, sub)
	sub, err = OnRequest(s.nc, MsgWorldListen, s.listenWorld)
	if err != nil {
		return fmt.Errorf("failed to subscribe to list actors requests: %w", err)
	}
	s.subs = append(s.subs, sub)
	sub, err = OnRequest(s.nc, MsgActorCreate, s.createActor)
	if err != nil {
		return fmt.Errorf("failed to subscribe to create actor requests: %w", err)
	}
	s.subs = append(s.subs, sub)
	sub, err = OnRequest(s.nc, MsgActorDelete, s.deleteActor)
	if err != nil {
		return fmt.Errorf("failed to subscribe to delete actor requests: %w", err)
	}
	s.subs = append(s.subs, sub)

	sub, err = OnRequest(s.nc, MsgPing, s.ping)
	if err != nil {
		return fmt.Errorf("failed to subscribe to ping requests: %w", err)
	}
	s.subs = append(s.subs, sub)

	sub, err = OnRequest(s.nc, MsgListActors, s.ListActors)
	if err != nil {
		return fmt.Errorf("failed to subscribe to list actors requests: %w", err)
	}
	s.subs = append(s.subs, sub)

	sub, err = OnRequest(s.nc, MsgCall, s.actorCall)
	if err != nil {
		return fmt.Errorf("failed to subscribe to actor call requests: %w", err)
	}
	s.subs = append(s.subs, sub)

	sub, err = OnRequest(s.nc, MsgWorldCall, s.worldCall)
	if err != nil {
		return fmt.Errorf("failed to subscribe to world call requests: %w", err)
	}
	s.subs = append(s.subs, sub)
	return nil
}

// Close closes the NATS connection
func (s *Service) Close() error {
	for _, unsub := range s.unsubs {
		unsub()
	}
	s.nc.Drain()
	return nil
}

func (s *Service) actorCall(msg *Msg) (*Msg, error) {
	var args []any
	err := json.Unmarshal(msg.Data, &args)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal actor call arguments: %w", err)
	}
	v, err := s.manager.CallActorMethod(msg.World, msg.Actor, msg.Member, args)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal actor call result: %w", err)
	}
	return &Msg{
		Type:  msg.Type,
		World: msg.World,
		Actor: msg.Actor,
		Data:  data,
	}, nil
}

func (s *Service) runScript(msg *Msg) (*Msg, error) {
	var script model.Script
	err := json.Unmarshal(msg.Data, &script)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal script: %w", err)
	}
	v, err := s.manager.RunScript(msg.World, script)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal script result: %w", err)
	}
	return &Msg{
		Type:  msg.Type,
		World: msg.World,
		Data:  data,
	}, nil
}

func (s *Service) getActorState(msg *Msg) (*Msg, error) {
	state := s.manager.GetActorState(msg.World, msg.Actor)
	data, err := json.Marshal(state)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal actor state: %w", err)
	}
	return &Msg{
		Type:  msg.Type,
		World: msg.World,
		Actor: msg.Actor,
		Data:  data,
	}, nil
}

func (s *Service) getActorValue(msg *Msg) (*Msg, error) {
	v := s.manager.GetActorValue(msg.World, msg.Actor, msg.Member)
	data, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal actor value: %w", err)
	}
	return &Msg{
		Type:  msg.Type,
		World: msg.World,
		Actor: msg.Actor,
		Data:  data,
	}, nil
}

func (s *Service) setActorValue(msg *Msg) (*Msg, error) {
	s.manager.SetActorValue(msg.World, msg.Actor, msg.Member, msg.Data)
	return &Msg{
		Type:  MsgSetValue,
		World: msg.World,
		Actor: msg.Actor,
	}, nil
}

func (s *Service) ListActors(msg *Msg) (*Msg, error) {
	actors := s.manager.ListActors(msg.World)
	data, err := json.Marshal(actors)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal actor list: %w", err)
	}
	return &Msg{
		Type:  msg.Type,
		World: msg.World,
		Data:  data,
	}, nil
}

func (s *Service) listenWorld(msg *Msg) (*Msg, error) {
	log.Info().Str("world", msg.World).Msg("listening to world")
	unsub := s.manager.WorldHooks(msg.World).Add(func(evt *model.SimEvent) {
		log.Info().Str("event", string(evt.Type)).Msg("world event")
		data, err := json.Marshal(evt)
		if err != nil {
			log.Error().Err(err).Msg("failed to marshal world event")
			return
		}
		msg := Msg{
			Type:  MsgWorldEvents,
			World: msg.World,
			Data:  data,
		}
		err = DoPublish(s.nc, &msg)
		if err != nil {
			log.Error().Err(err).Msg("failed to publish world event")
		}
	})
	s.unsubs = append(s.unsubs, unsub)
	return msg, nil
}

func (s *Service) createWorld(msg *Msg) (*Msg, error) {
	log.Info().Str("world", msg.World).Msg("creating world")
	var script model.Script
	err := json.Unmarshal(msg.Data, &script)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal script: %w", err)
	}
	if !script.IsEmpty() {
		v, err := s.manager.RunScript(msg.World, script)
		if err != nil {
			return nil, fmt.Errorf("failed to run initialization script: %w", err)
		}
		data, err := json.Marshal(v)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal initialization script result: %w", err)
		}
		return &Msg{
			Type:  msg.Type,
			World: msg.World,
			Data:  data,
		}, nil
	}
	s.manager.GetSimulation(msg.World)
	return &Msg{
		Type:  msg.Type,
		World: msg.World,
	}, nil
}

func (s *Service) deleteWorld(msg *Msg) (*Msg, error) {
	log.Info().Str("world", msg.World).Msg("deleting world")
	s.manager.DeleteSimulation(msg.World)
	return &Msg{
		Type:  msg.Type,
		World: msg.World,
	}, nil
}

func (s *Service) worldStatus(msg *Msg) (*Msg, error) {
	status := s.manager.GetSimulationStatus(msg.World)
	data, err := json.Marshal(status)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal world status: %w", err)
	}
	return &Msg{
		Type:  msg.Type,
		World: msg.World,
		Data:  data,
	}, nil
}

func (s *Service) createActor(msg *Msg) (*Msg, error) {
	log.Info().Str("world", msg.World).Str("actor", msg.Actor).Msg("creating actor")
	state := make(map[string]any)
	s.manager.CreateActor(msg.World, msg.Actor, state)
	return &Msg{
		Type:  msg.Type,
		World: msg.World,
		Actor: msg.Actor,
	}, nil
}

func (s *Service) deleteActor(msg *Msg) (*Msg, error) {
	log.Info().Str("world", msg.World).Str("actor", msg.Actor).Msg("deleting actor")
	s.manager.DeleteActor(msg.World, msg.Actor)
	return &Msg{
		Type:  msg.Type,
		World: msg.World,
		Actor: msg.Actor,
	}, nil
}

func (s *Service) ping(msg *Msg) (*Msg, error) {
	return &Msg{
		Type:  msg.Type,
		World: msg.World,
	}, nil
}

func (s *Service) worldCall(msg *Msg) (*Msg, error) {
	var args []any
	err := json.Unmarshal(msg.Data, &args)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal actor call arguments: %w", err)
	}
	log.Info().Str("world", msg.World).Str("member", msg.Member).Msg("calling world method")
	v, err := s.manager.CallWorldMethod(msg.World, msg.Member, args)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal actor call result: %w", err)
	}
	return &Msg{
		Type:  msg.Type,
		World: msg.World,
		Data:  data,
	}, nil
}
