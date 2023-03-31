package main

import (
	"log"
	"adventure-tools/subcommands/scaler"

	"github.com/alexflint/go-arg"
)

type Cli struct {
	Scale *scaler.Command `arg:"subcommand:scale" help:"use this command to scale png sprites into a resulting png"`
}

// Args is the command line arguments parsed as a struct. A subcommand being parsed as not nil
// means it has been invoked.
var Args Cli

func init() {
	arg.MustParse(&Args)
}

func main() {
	switch {
	case Args.Scale != nil:
		if err := Args.Scale.Run(); err != nil {
			log.Fatal(err)
		}
	}
}
