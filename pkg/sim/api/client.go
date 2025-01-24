package api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/sim/model"
	"github.com/nats-io/nats.go"
)

type ObjectAPI struct {
	world  string
	client *Client
}

func NewObjectAPI(world string, client *Client) *ObjectAPI {
	return &ObjectAPI{
		world:  world,
		client: client,
	}
}

func (o *ObjectAPI) GetProperties(name string) (map[string]any, error) {
	return o.client.GetActorState(o.world, name)
}

// SetProperties
func (o *ObjectAPI) SetProperties(name string, props map[string]any) error {
	return o.client.SetActorState(o.world, name, props)
}

func (o *ObjectAPI) GetProperty(name, member string) (any, error) {
	return o.client.GetActorValue(o.world, name, member)
}

// SetProperty sets a property on an object
func (o *ObjectAPI) SetProperty(name, member string, value any) error {
	return o.client.SetActorValue(o.world, name, member, value)
}

// InvokeMethod invokes a method on an object
func (o *ObjectAPI) InvokeMethod(name, member string, args []any) (any, error) {
	return o.client.ActorCall(o.world, name, member, args)
}

// EmitSignal emits a signal on an object
func (o *ObjectAPI) EmitSignal(name, signal string, args []any) error {
	return o.client.ActorEmitSignal(o.world, name, signal, args)
}

// Client provides a high-level API for interacting with the SimBus service
type Client struct {
	nc      *nats.Conn
	timeout time.Duration
}

// NewClient creates a new SimBus client
func NewClient(natsURL string) (*Client, error) {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}

	return &Client{
		nc:      nc,
		timeout: 30 * time.Second,
	}, nil
}

// Close closes the NATS connection
func (c *Client) Close() {
	c.nc.Close()
}

// SetTimeout sets the timeout for requests
func (c *Client) SetTimeout(timeout time.Duration) {
	c.timeout = timeout
}

func (c *Client) RunScript(world string, script model.Script) error {
	if world == "" {
		world = "demo"
	}
	log.Info().Str("world", world).Str("script", script.Name).Msg("client request run script")
	data, err := json.Marshal(script)
	if err != nil {
		return fmt.Errorf("failed to marshal script: %w", err)
	}
	req := Msg{
		Type:  MsgRunScript,
		World: world,
		Data:  data,
	}
	reply, err := DoRequest(c.nc, &req)
	if err != nil {
		return err
	}
	var result any
	err = json.Unmarshal(reply.Data, &result)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetActorState(world, name string) (map[string]any, error) {
	if world == "" {
		world = "demo"
	}
	log.Info().Str("world", world).Str("name", name).Msg("client.GetActorState")
	req := Msg{
		Type:  MsgGetState,
		World: world,
		Actor: name,
	}
	reply, err := DoRequest(c.nc, &req)
	if err != nil {
		return nil, err
	}
	var state map[string]any
	err = json.Unmarshal(reply.Data, &state)
	if err != nil {
		return nil, err
	}
	return state, nil
}

// GetActorValue retrieves a specific value from an actor
func (c *Client) GetActorValue(world, name, member string) (any, error) {
	log.Info().Str("world", world).Str("name", name).Str("member", member).Msg("client.GetActorValue")
	req := Msg{
		Type:   MsgGetValue,
		World:  world,
		Actor:  name,
		Member: member,
	}
	reply, err := DoRequest(c.nc, &req)
	if err != nil {
		return nil, err
	}
	var value any
	err = json.Unmarshal(reply.Data, &value)
	if err != nil {
		return nil, err
	}
	return value, nil
}

// SetActorValue sets a specific value on an actor
func (c *Client) SetActorValue(world, name, member string, value any) error {
	if world == "" {
		world = "demo"
	}
	log.Info().Str("world", world).Str("name", name).Str("member", member).Msg("client set actor value")
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value: %w", err)
	}
	req := Msg{
		Type:   MsgSetValue,
		World:  world,
		Actor:  name,
		Member: member,
		Data:   data,
	}
	_, err = DoRequest(c.nc, &req)
	return err
}

// WorldListen starts listening for events on a world
func (c *Client) WorldListen(world string) error {
	if world == "" {
		world = "demo"
	}
	log.Info().Str("world", world).Msg("client.WorldListen")
	req := Msg{
		Type:  MsgWorldListen,
		World: world,
	}
	_, err := DoRequest(c.nc, &req)
	return err
}

// ListActors lists the actors in a world
func (c *Client) ListActors(world string) ([]string, error) {
	if world == "" {
		world = "demo"
	}
	log.Info().Str("world", world).Msg("client.ListActors")
	req := Msg{
		Type:  MsgListActors,
		World: world,
	}
	reply, err := DoRequest(c.nc, &req)
	if err != nil {
		return nil, err
	}
	var actors []string
	err = json.Unmarshal(reply.Data, &actors)
	if err != nil {
		return nil, err
	}
	return actors, nil
}

// WorldClose closes a world
func (c *Client) WorldClose(world string) error {
	if world == "" {
		world = "demo"
	}
	log.Info().Str("world", world).Msg("client.WorldClose")
	req := Msg{
		Type:  MsgWorldClose,
		World: world,
	}
	_, err := DoRequest(c.nc, &req)
	return err
}

func (c *Client) ActorCall(world, actor, member string, args []any) (any, error) {
	if world == "" {
		world = "demo"
	}
	data, err := json.Marshal(args)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal arguments: %w", err)
	}
	req := Msg{
		Type:   MsgCall,
		World:  world,
		Actor:  actor,
		Member: member,
		Data:   data,
	}

	res, err := DoRequest(c.nc, &req)
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}

// WorldCallFunction calls a function on a world
func (c *Client) WorldCallFunction(world, name string, args []any) (any, error) {
	if world == "" {
		world = "demo"
	}
	log.Info().Str("world", world).Str("function", name).Msg("client request: call world function")
	data, err := json.Marshal(args)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal arguments: %w", err)
	}
	req := Msg{
		Type:   MsgWorldCall,
		World:  world,
		Member: name,
		Data:   data,
	}
	_, err = DoRequest(c.nc, &req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// CreateWorld creates a new world
func (c *Client) CreateWorld(name string) error {
	if name == "" {
		name = "demo"
	}
	log.Info().Str("name", name).Msg("client.CreateWorld")
	req := Msg{
		Type:  MsgWorldCreate,
		World: name,
	}
	_, err := DoRequest(c.nc, &req)
	return err
}

// DeleteWorld deletes a world
func (c *Client) DeleteWorld(world string) error {
	if world == "" {
		world = "demo"
	}
	log.Info().Str("world", world).Msg("client.DeleteWorld")
	req := Msg{
		Type:  MsgWorldDelete,
		World: world,
	}
	_, err := DoRequest(c.nc, &req)
	return err
}

// WorldStatus gets the status of a world
func (c *Client) WorldStatus(world string) (model.WorldStatus, error) {
	if world == "" {
		world = "demo"
	}
	log.Info().Str("world", world).Msg("client.WorldStatus")
	req := Msg{
		Type:  MsgWorldStatus,
		World: world,
	}
	res, err := DoRequest(c.nc, &req)
	if err != nil {
		return model.WorldStatus{}, err
	}
	var status model.WorldStatus
	err = json.Unmarshal(res.Data, &status)
	if err != nil {
		return model.WorldStatus{}, err
	}
	return status, nil
}

// OnWorldEvents registers a callback for world events
func (c *Client) OnWorldEvents(fn func(evt *model.SimEvent) error) (*nats.Subscription, error) {

	return OnPublish(c.nc, MsgWorldEvents, func(msg *Msg) error {
		var event model.SimEvent
		err := json.Unmarshal(msg.Data, &event)
		if err != nil {
			return err
		}
		return fn(&event)
	})
}

func (c *Client) CreateActor(world, name string) error {
	if world == "" {
		world = "demo"
	}
	log.Info().Str("world", world).Str("name", name).Msg("client.CreateActor")
	req := Msg{
		Type:  MsgActorCreate,
		World: world,
		Actor: name,
	}
	_, err := DoRequest(c.nc, &req)
	return err
}

func (c *Client) DeleteActor(world, name string) error {
	if world == "" {
		world = "demo"
	}
	log.Info().Str("world", world).Str("name", name).Msg("client.DeleteActor")
	req := Msg{
		Type:  MsgActorDelete,
		World: world,
		Actor: name,
	}
	_, err := DoRequest(c.nc, &req)
	return err
}

// SetActorState sets the state of an actor
func (c *Client) SetActorState(world, name string, state map[string]any) error {
	if world == "" {
		world = "demo"
	}
	log.Info().Str("world", world).Str("name", name).Msg("client.SetActorState")
	data, err := json.Marshal(state)
	if err != nil {
		return fmt.Errorf("failed to marshal state: %w", err)
	}
	req := Msg{
		Type:  MsgSetState,
		World: world,
		Actor: name,
		Data:  data,
	}
	_, err = DoRequest(c.nc, &req)
	return err
}

func (c *Client) ActorEmitSignal(world, name, signal string, args []any) error {
	if world == "" {
		world = "demo"
	}
	log.Info().Str("world", world).Str("name", name).Str("signal", signal).Msg("client.ActorEmitSignal")
	data, err := json.Marshal(args)
	if err != nil {
		return fmt.Errorf("failed to marshal arguments: %w", err)
	}
	req := Msg{
		Type:   MsgSignal,
		World:  world,
		Actor:  name,
		Member: signal,
		Data:   data,
	}
	return DoPublish(c.nc, &req)
}

// Ping sends a ping to the server
func (c *Client) Ping() error {
	req := Msg{
		Type: MsgPing,
	}
	_, err := DoRequest(c.nc, &req)
	return err
}

func (c *Client) ObjectAPI(world string) *ObjectAPI {
	return NewObjectAPI(world, c)
}
