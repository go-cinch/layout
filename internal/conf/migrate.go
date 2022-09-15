package conf

import "embed"

//go:embed db/*.sql
var SqlFiles embed.FS
