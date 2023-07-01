package server

import (
	"embed"

	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer)

//go:embed middleware/locales
var locales embed.FS
