// Olink Adapter for Simulation
//
// This package provides an adapter for simulation to be used with object-link.
// ObjectLink is a websocket based protocol for communication between client and server.
// A Hub on the server side is used to manage all connections and objects.
// ON each websocket connection a remote node is created.
// On a link message the node will be linked to the object.
// For this the object must be registered as a generic source.
// In case a source is not found the node will be linked to a default object.
// We set the default source for the registry to a generic object.
package olnk
