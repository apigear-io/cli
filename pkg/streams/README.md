# ApiGear Streams

## Concepts

This package provides the streaming services for ApiGear, built on top of NATS and NATS JetStream. 

- A stream is a sequence of messages, and a subject is a string that identifies a stream. 
- Messages are published to subjects, and subscribers can receive messages from subjects. 
- A store is a KV store that persists state for a specific domain.
- A buffer is a temporary storage for messages that are not yet processed by subscribers. 
- A consumer is an entity that subscribes to a subject and processes messages from that subject.
- A server is a NATS server instance that manages the streams, subjects, and messages.
- A monitor is a http endpoint that ingests messages from devices and forwards them to the appropriate subjects.

```
http -> monitor -> subject (deviceID) -> buffer (5min) -> store (recordingID) -> consumer (replay)
```

### Server

Initially we launch a NATS server with JetStream enabled. The server manages the streams, subjects, and messages. The server has an own disk storage for persistent streams, and an in-memory storage for temporary streams.
The service also loads the controller service and the buffer service, which are used to manage the streams and buffers.

### Stream Monitor

We then attach a HTTP monitor to the server, which listens for incoming messages from devices. The monitor forwards the messages to the appropriate subjects in the NATS server, based on the device ID. These messages are not persisted, but can be processed by subscribers in real-time. Device information is stored in the device KV store.

### Buffer Window

To allow later to record messages form the past, we can attach a buffer to the device ID subject. The buffer stores messages temporarily until they are processed by subscribers, or the retention policy (e.g. 5min) deletes the messages. This allows us to record messages from the past, even if the subscribers were not connected at the time the messages were published. The buffer informatin is attached to the device KV store.

### Recording

We can start the recording of a device monitoring data, which creates a persistent stream for the device ID subject. The stream stores all messages published to the subject, and allows subscribers to receive messages from the past. Recording state is tracked in the session KV store. The stream can be configured with a retention policy (e.g. 24h), which deletes messages older than the specified duration. A recording has a session ID, which is used to identify the recording session for later retrieval.

### Replay 

We can then replay the recorded messages from a specific recording session. This creates a consumer that subscribes to the device ID subject, and processes messages from the persistent stream. The consumer can be configured with a start time, which allows to replay messages from a specific point in time and a speed factor, to replay messages faster or slower than real-time. The consumer processes messages in the order they were published, and can be stopped and restarted as needed.

## CLI Usage

```
# start a NATS server with JetStream enabled
apigear streams serve

# generate sample messages for a device
apigear streams data generate -c 1000 -o data.jsonl -t examples/orders.tmpl

# send sample messages to the monitor stream endpoint
apigear streams data publish --device-id 1234 --file data.jsonl --intervall 1s

# display the monitor stream messages in real-time
apigear streams data tail --device-id 1234

