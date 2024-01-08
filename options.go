package decade

import "github.com/urfave/cli/v2"

type Option interface {
	apply(*Decade)
}

type driverOptionFunc func(*Decade)

func (f driverOptionFunc) apply(d *Decade) {
	f(d)
}

func AddFlag(flag cli.Flag) Option {
	return driverOptionFunc(func(dd *Decade) {
		dd.cli.Flags = append(dd.cli.Flags, flag)
	})
}

func AddCommand(command *cli.Command) Option {
	return driverOptionFunc(func(dd *Decade) {
		dd.cli.Commands = append(dd.cli.Commands, command)
	})
}

func WithAfterFunc(f cli.AfterFunc) Option {
	return driverOptionFunc(func(dd *Decade) {
		dd.cli.After = f
	})
}
