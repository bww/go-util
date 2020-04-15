package main

import (
	"flag"
)

type globalFlags struct {
	Secret  *string
	Salt    *string
	Debug   *bool
	Verbose *bool
	Trace   *bool
}

func bindGlobalFlags(cmdline *flag.FlagSet) globalFlags {
	return globalFlags{
		Secret:  cmdline.String("secret", "", "The secret to use when deriving keys for message validation."),
		Salt:    cmdline.String("salt", "", "The salt to use when deriving keys for message validation."),
		Debug:   cmdline.Bool("debug", false, "Enable debugging mode."),
		Verbose: cmdline.Bool("verbose", false, "Enable verbose debugging mode."),
		Trace:   cmdline.Bool("trace", false, "Enable tracing mode."),
	}
}
