package streams

import (
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/nats-io/nats.go"
)

func PublishMonitorMessage(nc *nats.Conn, deviceId string, data []byte) error {
	log.Debug().Msgf("publish monitor message from device %s", deviceId)
	return nc.Publish(config.DeviceSubject(config.MonitorSubject, deviceId), data)
}
