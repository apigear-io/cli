# API Simulation

The API simulation is a tool that can be used to simulate the API calls and the performance of the API. The simulation listens on a websocket and shows the requests and the performance of the API.

The simulation is stored in either a JSON (YAML) file or run via a JavaScript file.

## Usage

Start the simulation with a JS scenario file:

```sh
$ cli simulation --port 8080 --scenario ./behavior.js
```

Use YAML (JSON) base API behavior

```sh
$ cli simulation --port 8080 --scenario ./behavior.yaml
```

Send scripted API events to the simulation

```sh
$ cli simulation --port 8080 --send ./test/events.ndjson
$ cli simulation --port 8080 --send ./test/events.csv
```
