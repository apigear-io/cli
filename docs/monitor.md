# API Monitor

The API monitor is a tool that can be used to monitor the API calls and the performance of the API. The monitor listens on a HTTP port and shows the requests and the performance of the API.

## Usage

Start the monitor

```sh
$ cli monitor --port 8080
```

Send scripted API events to the monitor

```sh
$ cli monitor --port 8080 --script ./test/events.ndjson
$ cli monitor --port 8080 --script ./test/events.csv
```

- `--interval`: The interval in seconds to send the events.
- `--repeat`: The number of times to send the events.
