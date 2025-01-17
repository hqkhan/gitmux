package main

import (
	_ "embed"
	"fmt"

	"github.com/arl/gitmux/tmux"
	"gopkg.in/yaml.v3"
)

// Config configures output formatting.
type Config struct{ Tmux tmux.Config }

// default config (decoded in init)
var defaultCfg Config

//go:embed .gitmux.yml
var cfgBytes []byte

func init() {
	if err := yaml.Unmarshal(cfgBytes, &defaultCfg); err != nil {
		panic(fmt.Sprintf("default config is invalid: %v", err))
	}
}
