# Concepts

ApiGear is a code generator based on the object api. The API describes a set of interfaces organized inside modules. The code transformation is done using templates and language specific filters.

Interface transformation can result into a simple local API or in more complex use cases result into a remote API.

For API development ApiGear CLI provides also an API event monitor as also an API simulation. The monitor traces all API events and the simulation allows to test the API without the need of a API backend.

## Code Generation

## API Monitoring

The API monitor listens on http://127.0.0.1:5555/monitor/{deviceId}/. Events come in as HTTP post requests, where the body is an event structure.

The event structure is a json object with the following fields:

```json
{
  "kind": "string",
  "deviceId": "id",
  "source": "id",
  "timestamp": "time",
  "symbol": "string",
  "params": [],
  "props": {}
}
```

## API Simulation
