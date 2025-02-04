# Simulation Commands

## apigear sim run demo.js

sends the demo.js script to the simulation server and runs it.
Prints out the world id.

## apigear sim stop <world id>

stops the simulation server with the given world id.


## apigear sim start <world id>

## apigear sim inspect <world id>

## apigear sim call <world id> <method> <args>

## apigear sim set world <world id>

# shows the current state of the simulation server with the given world id.


## sim server

Run simulation server using NATS as also a simulation olink server and the http server for API monitoring.

@TODO: move server to own subcommand

## apigear server run

Run apigear server using NATS as also a simulation olink server and the http server for API monitoring.

