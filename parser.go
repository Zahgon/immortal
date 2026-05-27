package immortal

import (
	"flag"
	"os/user"
)

// Parser interface
type Parser interface {
	Parse(fs *flag.FlagSet) (*Flags, error)
	parseYml(file string) (*Config, error)
	checkWrkdir(dir string) error
	parseEnvdir(dir string) (map[string]string, error)
	checkUser(username string) (*user.User, error)
}

// Parse implements parser
type Parse struct {
	Flags
	UserLookup func(username string) (*user.User, error)
}

// Parse parse the command line flags
func (p *Parse) Parse(fs *flag.FlagSet) (*Flags, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *Parse) parseYml(file string) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

// set defaults

func (p *Parse) checkWrkdir(dir string) error { _ = "STUB: not implemented"; return nil }

func (p *Parse) parseEnvdir(dir string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// checkUser needs cgo
func (p *Parse) checkUser(u string) (*user.User, error) { _ = "STUB: not implemented"; return nil, nil }

// Usage prints to standard error a usage message
func (p *Parse) Usage(fs *flag.FlagSet) func() { _ = "STUB: not implemented"; return nil }

// ParseArgs parse command arguments
func ParseArgs(p Parser, fs *flag.FlagSet) (cfg *Config, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if -v

// if -ctl, defaults to /var/run/immortal

// if -cc

// if -c

// parse the `run.yml` file

// Save ConfigFile to use service name derived from file name instead of the PID

// Cmd is mandatory, is the command that needs to be supervised

// split command into a slice of strings

// Change working directory, will: cd Cwd before starting

// The user to run the process on behalf

// print config and exit 0 if config ok

// if no args

// create new cfg if not using run.yml

// if -d

// if -e

// if -f

// if -l

// if -logger

// if -P

// if -p

// if -r

// if -w

// if -u
