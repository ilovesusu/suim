package server

import "github.com/google/wire"

// ProviderSet is server providers.
//var ProviderSet = wire.NewSet(NewTCPServer, NewWebSocketServer, NewEtcdServer)
var ProviderSet = wire.NewSet(NewTCPServer, NewWebSocketServer)
