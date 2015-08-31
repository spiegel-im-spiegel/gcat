package main

import (
	"os"

	"github.com/spiegel-im-spiegel/gutil"
	"gcat/internal/facade"
)

func main() {
	cli := &gutil.CliContext{Reader: os.Stdin, Writer: os.Stdout, ErrorWriter: os.Stderr}
	facadeCxt := &facade.Context{Cli: cli, CommandName: Name, Version: Version}
	os.Exit(facadeCxt.Run(os.Args))
}