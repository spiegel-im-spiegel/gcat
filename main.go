package main

import (
	"os"

	"github.com/spiegel-im-spiegel/gutil"
	"github.com/spiegel-im-spiegel/gcat/internal/facade"
)

func main() {
	cli := &gutil.CliUi{Reader: os.Stdin, Writer: os.Stdout, ErrorWriter: os.Stderr}
	facadeCxt := &facade.Context{Cli: cli, CommandName: Name, Version: Version}
	os.Exit(facadeCxt.Run(os.Args))
}
