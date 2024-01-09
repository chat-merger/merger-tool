package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type Config struct {
	//Version  string
	ChatsDir   string
	ServerPort int
}

// Flag-feature part:

// FlagSet is Config factory
type FlagSet struct {
	cfg Config
	fs  *flag.FlagSet
}

func InitFlagSet() *FlagSet {
	cfgFs := new(FlagSet)
	cfgFs.fs = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	cfgFs.fs.StringVar(&cfgFs.cfg.ChatsDir, flagChatsDir, "", "path to directory with chats")
	cfgFs.fs.IntVar(&cfgFs.cfg.ServerPort, flagServerPort, 0, "run http server on this port")
	return cfgFs
}

// cleanLastCfg clean parsed values
func (c *FlagSet) cleanLastCfg() {
	c.cfg.ChatsDir = ""
	c.cfg.ServerPort = 0
}

// Flag names:

const (
	flagChatsDir   = "chats-dir"
	flagServerPort = "port"
)

// Usage printing "how usage flags" message
func (c *FlagSet) Usage() { c.fs.Usage() }

// Parse is Config factory method
func (c *FlagSet) Parse(args []string) (*Config, error) {
	missingArgExit := func(argName string) error {
		return fmt.Errorf("missing `%s` argument: %w", argName, WrongArgumentError)
	}

	err := c.fs.Parse(args)
	if err != nil {
		return nil, fmt.Errorf("parse given config arguments: %w", err)
	}
	newCfg := c.cfg // copy parsed values
	c.cleanLastCfg()

	if newCfg.ChatsDir == "" {
		return nil, missingArgExit(flagChatsDir)
	}
	if newCfg.ServerPort == 0 {
		return nil, missingArgExit(flagServerPort)
	}

	return &newCfg, nil
}

var WrongArgumentError = errors.New("wrong argument")
