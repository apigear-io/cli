# Streams


## Server startup

To start the API Gear server with streams support, run the following command:
```sh
apigear serve
```

## Recording a device stream

### Producing monitoring data

To simulate a device producing monitoring data, you can publish data from a NDJSON file:

```sh
apigear stream publish --file data/mon/sample.ndjson --device 123 --interval 1s
```

This will send every 1s a new line from the `sample.ndjson` file to the monitoring system for device `123`.

### Starting the recording

To start recording a device stream, use the following command:

```sh
apigear stream record --device 123
```

This will output:
```
recording started session=<session-id>
```

You will need this session ID to stop the recording and play it back later.

### Stopping the recording

To stop the recording, use the following command with the session ID you received earlier:
```sh
apigear stream stop --session <session-id>
```

### Viewing recorded sessions

List all recorded sessions:
```sh
apigear stream ls
```

Show detailed information about a specific session:
```sh
apigear stream show --session <session-id>
```

## Playing back a recorded stream

### Olink connection

To play back the recorded stream, first connect to the Olink server and link the desired object:
```sh
apigear olink
> connect
> link demo.Counter
```

### Playback command

Then, use the following command to play back the recorded stream using the session ID:
```sh
apigear stream play --session <session-id>
```

You can control playback speed with the `--speed` flag (e.g., `--speed 0.25` for quarter speed, `--speed 2` for double speed).

## Additional commands

### Monitoring live streams

To monitor live data from a device in real-time:
```sh
apigear stream tail --device 123
```

### Generating test data

To generate test monitoring data from a template:
```sh
apigear stream generate --template template.json --output test-data.ndjson --count 1000
```

### Managing devices

Set device metadata:
```sh
apigear stream device-set --device 123 --desc "Test Device" --location "Lab A" --owner "Team X"
```

Get device information:
```sh
apigear stream device-get --device 123
```

List all devices:
```sh
apigear stream device-ls
```

### Managing device buffers

Enable buffering for a device (useful with `--pre-roll` during recording):
```sh
apigear stream device buffer enable --device 123 --window 5m
```

Get buffer information:
```sh
apigear stream device buffer info --device 123
```

## Behind the scenes

- When you start the apigear server, you actually start a NATS server with JetStream enabled.
- When you publish data to the monitoring system using `stream publish`, the data is sent to a NATS subject based on the device ID (e.g., `monitor.123`).
- To record this data, run `apigear stream record --device 123`, which creates a JetStream consumer subscription to the monitoring subject for the specified device ID.
- The recording entry and state are stored in a KV store in JetStream. This allows us to watch the state and resume interrupted recordings.
- Device information is also stored in the KV store, so we know which device each recording belongs to.
- The recording subscription stores the recorded data in a JetStream stream under a unique session ID.
- To stop the recording, run `apigear stream stop --session <session-id>`, which stops the subscription and finalizes the recorded data.
- Before playback, connect to the Olink server and link the desired object (e.g., `demo.Counter`) to receive the playback data.
- Finally, run `apigear stream play --session <session-id>`, which reads the recorded data from JetStream and publishes it to the linked Olink object using a JetStream consumer.
- You can control the playback speed using the `--speed` flag (e.g., `--speed 2` for double speed).
- Device buffers can be enabled with `--pre-roll` during recording to capture data from before the recording started.
