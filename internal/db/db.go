package db

import "embed"

const SQLRoot = "migrations"

//go:embed migrations/*.sql
var SQLFiles embed.FS
